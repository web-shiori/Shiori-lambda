package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

func s3Handler(ctx context.Context, event events.S3Event) {
	fmt.Println("Lambda Start")
	for _, record := range event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key
		fmt.Printf("bucket: %s, key: %s\n", bucket, key)
		pageNum, err := detectPageNumber(bucket, key)
		if err != nil {
			panic(err)
		}
		fmt.Printf("pageNum: %d\n", pageNum)
	}
	fmt.Println("Lambda Finish")
}

// PDFのページ数を取得する
func detectPageNumber(bucket string, key string) (pageNum int, err error) {
	var textractSession *textract.Textract
	var simplePageNumExtracter SimplePageNumExtracter
	fmt.Println(bucket)

	// S3のファイルをTextractにかける
	textractSession = textract.New(session.Must(session.NewSession(&aws.Config{
		// TODO: 環境変数から取得する
		Region: aws.String(os.Getenv("textractRegionName")),
	})))
	resp, err := textractSession.DetectDocumentText(&textract.DetectDocumentTextInput{
		Document: &textract.Document{
			S3Object: &textract.S3Object{
				Bucket: aws.String(bucket),
				Name:   aws.String(key),
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	fmt.Println(resp)

	detectWordSlice, err := detectDocumentTextOutputToStringSlice(resp)
	fmt.Println(detectWordSlice)

	// PDFのページ数を抽出する
	pageNum, err = simplePageNumExtracter.extractPageNum(detectWordSlice)
	if err != nil {
		return 0, err
	}

	return
}

// Textractで検知した文字列をスライスにする
func detectDocumentTextOutputToStringSlice(textOutput *textract.DetectDocumentTextOutput) (detectWordSlice []string, err error) {
	for _, w := range textOutput.Blocks {
		if *w.BlockType == "WORD" {
			detectWordSlice = append(detectWordSlice, *w.Text)
		}
	}
	return
}

func main() {
	lambda.Start(s3Handler)

	// ローカル開発用
	//detectPageNumber("web-snapshot-s3-us-east-1", "sample.png")
}
