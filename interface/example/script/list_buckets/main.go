package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String("http://localhost:4572"),
		Region:           aws.String("ap-northeast1"),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		log.Fatal(err)
	}
	svc := s3.New(sess)
	output, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result of List Bucket")
	for _, b := range output.Buckets {
		if b != nil && b.Name != nil {
			fmt.Println(*b.Name)
		}
	}
}
