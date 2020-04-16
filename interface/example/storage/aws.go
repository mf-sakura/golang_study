package storage

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// AWSConfig is Config for aws
type AWSConfig struct {
	EndpointURL string
	BucketName  string
	Region      string
}

// S3Client is storage client implementation for aws s3
type S3Client struct {
	Client     *s3.S3
	BucketName string
}

// NewAWSClient returns new AWSClient
func NewAWSClient(config *AWSConfig) (Client, error) {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(config.EndpointURL),
		Region:   aws.String(config.Region),
		// localstackを利用する為に必要
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)
	return &S3Client{
		Client:     svc,
		BucketName: config.BucketName,
	}, nil
}

// PutObject copies source file to destination
func (c *S3Client) PutObject(sourcePath, destinationPath string) error {
	file, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = c.Client.PutObject(&s3.PutObjectInput{
		// *stringを渡す必要があるので、AWSの関数を使って変換する
		Bucket: aws.String(c.BucketName),
		Key:    aws.String(destinationPath),
		Body:   file,
	})
	if err != nil {
		return err
	}
	return nil
}

// GetObject return files based on objectPath
func (c S3Client) GetObject(objectPath string) (io.ReadCloser, error) {
	output, err := c.Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.BucketName),
		Key:    aws.String(objectPath),
	})
	if err != nil {
		return nil, err
	}
	return output.Body, nil
}
