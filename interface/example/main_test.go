package main

import (
	"io"
	"strings"
	"testing"

	"github.com/mf-sakura/golang_study/interface/example/storage"
	"github.com/pkg/errors"
)

// mockStorageClient is storage client implementation for testing
type mockStorageClient struct {
	isError     bool
	fileContent string
}
type mockIOReadCloser struct {
	io.Reader
}

func (rc *mockIOReadCloser) Close() error {
	return nil
}
func newMockStorageClient(isError bool, content string) storage.Client {
	return &mockStorageClient{
		isError:     isError,
		fileContent: content,
	}
}
func (c *mockStorageClient) PutObject(sourcePath, destinationPath string) error {
	if c.isError {
		return errors.New("error")
	}
	return nil
}
func (c *mockStorageClient) GetObject(objectPath string) (io.ReadCloser, error) {
	if c.isError {
		return nil, errors.New("error")
	}
	return &mockIOReadCloser{strings.NewReader(c.fileContent)}, nil

}

func TestReadFile(t *testing.T) {
	type args struct {
		c storage.Client
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{c: newMockStorageClient(false, "abc")},
			want: "abc",
		},
		{
			name:    "異常系 Client.GetObjectがエラー",
			args:    args{c: newMockStorageClient(true, "")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
