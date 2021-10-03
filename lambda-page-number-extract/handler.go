package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

type OCRClient interface {
	DetectDocumentText(input *textract.DetectDocumentTextInput) (*textract.DetectDocumentTextOutput, error)
}

// PDFのページ数を取得する
func detectPageNumber(c OCRClient, pageNumExtractor PageNumExtractor, bucket string, key string) (pageNum int, err error) {
	// S3のファイルをOCRにかける
	input := &textract.DetectDocumentTextInput{
		Document: &textract.Document{
			S3Object: &textract.S3Object{
				Bucket: aws.String(bucket),
				Name:   aws.String(key),
			},
		},
	}
	resp, err := c.DetectDocumentText(input)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	detectWordSlice, err := detectDocumentTextOutputToStringSlice(resp)
	fmt.Println(detectWordSlice)

	// PDFのページ数を抽出する
	pageNum, err = pageNumExtractor.extractPageNum(detectWordSlice)
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

func s3Handler(ctx context.Context, event events.S3Event) {
	fmt.Println("Lambda Start")
	for _, record := range event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key
		fmt.Printf("bucket: %s, key: %s\n", bucket, key)

		region := os.Getenv("textractRegionName")
		textractSession := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region),
		}))
		textractClient := textract.New(textractSession)
		simplePageNumExtractor := new(SimplePageNumExtractor)
		pageNum, err := detectPageNumber(textractClient, simplePageNumExtractor, bucket, key)
		if err != nil {
			panic(err)
		}
		fmt.Printf("pageNum: %d\n", pageNum)
	}
	fmt.Println("Lambda Finish")
}

func main() {
	lambda.Start(s3Handler)

	// NOTE: ローカル開発用
	//bucket := "web-snapshot-s3-us-east-1"
	//key := "sample.png"
	//region := os.Getenv("textractRegionName")
	//textractSession := session.Must(session.NewSession(&aws.Config{
	//	Region: aws.String(region),
	//}))
	//textractClient := textract.New(textractSession)
	//simplePageNumExtractor := new(SimplePageNumExtractor)
	//detectPageNumber(textractClient, simplePageNumExtractor, bucket, key)
}
