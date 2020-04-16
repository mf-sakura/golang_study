package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"

	"github.com/mf-sakura/golang_study/interface/example/storage"
	"github.com/pkg/errors"
)

const (
	savedFolder = "output"
)

var (
	inputFilePath string
	savedFileName string
	action        string
	provider      string
)

func main() {

	flag.StringVar(&inputFilePath, "f", "input/test.txt", "Path of input file")
	flag.StringVar(&savedFileName, "s", "output/test.txt", "Name of saved file")
	flag.StringVar(&action, "a", "put", "action")
	flag.StringVar(&provider, "p", "provider", "aws or local")
	flag.Parse()

	c, err := NewStorageClient()
	if err != nil {
		log.Fatalf("NewStorageClient failed. %v", err)
	}
	switch action {
	case "write":
		if err := WriteFile(c); err != nil {
			log.Fatalf("PutFile failed. %v", err)
		}
	case "read":
		content, err := ReadFile(c)
		if err != nil {
			log.Fatalf("ReadFile failed. %v", err)
		}
		fmt.Println(content)
	default:
		log.Fatal("unrecognized action")
	}

}

// NewStorageClient returns storage.Client based on provider
func NewStorageClient() (storage.Client, error) {
	switch provider {
	case "aws":
		return storage.NewAWSClient(&storage.AWSConfig{
			BucketName:  "test",
			EndpointURL: "http://localhost:4572",
			Region:      "ap-northeast-1",
		})
	case "local":
		return storage.NewLocalIOClient(), nil
	default:
		return nil, errors.New("unrecognized provider")
	}
}

// WriteFile write file to storage
func WriteFile(c storage.Client) error {
	return c.PutObject(inputFilePath, fmt.Sprintf("%s/%s", savedFolder, savedFileName))
}

// ReadFile returns file content
func ReadFile(c storage.Client) (string, error) {
	readCloser, err := c.GetObject(fmt.Sprintf("%s/%s", savedFolder, savedFileName))
	if err != nil {
		return "", err
	}
	defer readCloser.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(readCloser); err != nil {
		return "", err
	}

	return buf.String(), nil
}
