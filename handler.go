package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func s3Handler(ctx context.Context, event events.S3Event) (err error) {
	for _, record := range event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key
		fmt.Printf("bucket: %s, key: %s", bucket, key)
	}
	return err
}

func main() {
	lambda.Start(s3Handler)
}
