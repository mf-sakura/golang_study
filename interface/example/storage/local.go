package storage

import (
	"io"
	"os"
)

// LocalIOClient is storage client implementation for local storage
type LocalIOClient struct{}

// NewLocalIOClient returns new LocalIOClient
func NewLocalIOClient() Client {
	return &LocalIOClient{}
}

// PutObject copies source file to destination
func (c *LocalIOClient) PutObject(sourcePath, destinationPath string) error {

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
}

// GetObject return files based on objectPath
func (c *LocalIOClient) GetObject(objectPath string) (io.ReadCloser, error) {
	return os.Open(objectPath)
}
