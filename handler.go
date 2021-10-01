package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
	"io/ioutil"
)

func s3Handler(ctx context.Context, event events.S3Event) {
	fmt.Println("Lambda Start")
	for _, record := range event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key
		fmt.Printf("bucket: %s, key: %s\n", bucket, key)
	}
	fmt.Println("Lambda Finish")
}

var textractSession *textract.Textract

func main() {
	//lambda.Start(s3Handler)

	// TODO: S3のファイルを参照するようにする
	file, err := ioutil.ReadFile("sample.png")
	if err != nil {
		panic(err)
	}

	// TODO: refactoring
	textractSession = textract.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})))

	resp, err := textractSession.DetectDocumentText(&textract.DetectDocumentTextInput{
		Document: &textract.Document{
			Bytes: file,
		},
	})
	if err != nil {
		panic(err)
	}

	for _, w := range resp.Blocks {
		if *w.BlockType == "WORD" {
			fmt.Println(*w.Text)
		}
	}
}
