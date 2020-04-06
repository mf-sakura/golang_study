package storage

import "io"

// Client is interface of client to access to strage
type Client interface {
	PutObject(sourcePath, destinationPath string) error
	GetObject(objectPath string) (io.ReadCloser, error)
}
