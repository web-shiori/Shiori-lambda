package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func s3Handler(ctx context.Context, event events.S3Event) {
	if len(event.Records) != 1 {
		panic(fmt.Errorf("Error length of event.Records is not 1. "))
	}
	r := event.Records[0]

	// PDFのページ数を取得する.
	//extractPDFPageNum(r)

	// ページ数をPUTリクエストする.
	contentID, err := getContentID(r.S3.Object)
	if err != nil {
		panic(err)
	}
	err = putPDFPageNum(contentID, 10)
	if err != nil {
		panic(err)
	}

	// PDFのスクリーンショットを削除する.
}

func main() {
	fmt.Println("------Lambda Start------")
	lambda.Start(s3Handler)
	defer fmt.Println("------Lambda Finish------")

	// NOTE: ローカル開発用
	//bucket := "web-snapshot-s3-us-east-1"
	//key := "examples.png"
	//region := os.Getenv("textractRegionName")
	//textractSession := session.Must(session.NewSession(&aws.Config{
	//	Region: aws.String(region),
	//}))
	//textractClient := textract.New(textractSession)
	//simplePageNumExtractor := new(SimplePageNumExtractor)
	//detectPageNumber(textractClient, simplePageNumExtractor, bucket, key)
}
