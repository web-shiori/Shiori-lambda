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
	s := s3Service{record: r}

	// PDFのページ数を取得する.
	pageNum, err := extractPDFPageNum(&s)
	if err != nil {
		panic(err)
	}

	// ページ数をPUTリクエストする.
	contentID, err := s.getContentID()
	if err != nil {
		panic(err)
	}
	fmt.Printf("------contentID: %s\n", contentID)
	err = putPDFPageNum(contentID, pageNum)
	if err != nil {
		panic(err)
	}

	// PDFのスクリーンショットを削除する.
	fmt.Println("一旦削除はパス")
	//fmt.Println("------Delete object------")
	//err = s.deleteObject()
	//if err != nil {
	//	fmt.Println("delete object error")
	//	panic(err)
	//}
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
