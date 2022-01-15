package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func s3Handler(ctx context.Context, event events.S3Event) {
	err := ImageColorInvert(event)
	if err != nil {
		panic(err)
	}
	fmt.Println("------Lambda invert-image-color Finish------")
}

func main() {
	fmt.Println("------Lambda invert-image-color Start------")
	lambda.Start(s3Handler)
}
