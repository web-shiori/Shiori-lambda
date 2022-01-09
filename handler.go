package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// PDFのスクリーンショットを削除

func s3Handler(ctx context.Context, event events.S3Event) {
	//extractPDFPageNum(event.Records[0])
	putPDFPageNum(10)
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
