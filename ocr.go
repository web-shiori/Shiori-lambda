package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

type OCRClient interface {
	DetectDocumentText(input *textract.DetectDocumentTextInput) (*textract.DetectDocumentTextOutput, error)
}

// PDFのページ数を取得する
func detectPageNumber(c OCRClient, pageNumExtractor PageNumExtractor, s3Object *textract.S3Object) (pageNum int, err error) {
	// S3のファイルをOCRにかける
	input := &textract.DetectDocumentTextInput{
		Document: &textract.Document{
			S3Object: s3Object,
		},
	}
	resp, err := c.DetectDocumentText(input)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	detectWordSlice, err := detectDocumentTextOutputToStringSlice(resp)
	if err != nil {
		return 0, err
	}
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
func extractPDFPageNum(service *s3Service) (int, error) {
	fmt.Println("------PDFのページ数を取得------")
	fmt.Printf("bucket: %s\n", service.record.S3.Bucket.Name)
	fmt.Printf("key: %s\n", service.record.S3.Object.Key)

	region := os.Getenv("textractRegionName")
	textractSession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	textractClient := textract.New(textractSession)
	simplePageNumExtractor := new(SimplePageNumExtractor)
	s3Object := service.getTextractS3Object()
	pageNum, err := detectPageNumber(textractClient, simplePageNumExtractor, s3Object)
	if err != nil {
		return 0, err
	}
	fmt.Printf("ページ数: %d\n", pageNum)
	return pageNum, nil
}
