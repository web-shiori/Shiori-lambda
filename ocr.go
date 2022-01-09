package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
	"os"
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

// PDFのページ数を取得
func extractPDFPageNum(record events.S3EventRecord) {
	fmt.Println("------PDFのページ数を取得------")
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
	fmt.Printf("ページ数: %d\n", pageNum)
}
