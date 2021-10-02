package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
		pageNumber, err := detectPageNumber(bucket, key)
		if err != nil {
			panic(err)
		}
		fmt.Println(pageNumber)
	}
	fmt.Println("Lambda Finish")
}

// PDFのページ数を取得する
func detectPageNumber(bucket string, key string) (pageNum int, err error) {
	var textractSession *textract.Textract

	// S3のファイルをTextractにかける
	textractSession = textract.New(session.Must(session.NewSession(&aws.Config{
		// TODO: 環境変数から取得する
		Region: aws.String("us-east-1"),
	})))

	resp, err := textractSession.DetectDocumentText(&textract.DetectDocumentTextInput{
		Document: &textract.Document{
			S3Object: &textract.S3Object{
				Bucket:  aws.String(bucket),
				Name:    aws.String(key),
			},
		},
	})
	if err != nil {
		return 0, err
	}


	// PDFのページ数を抽出する
	// TODO: impl
	/*
	レスポンスの例
	AIDataTechnologyMap_210520.po
	6
	/
	108
	I
	-
	100%
	+
	I
	:
	a
	am
	4
	Anaheim
	44
	Corona
	46
	Zumwalt
	48
	Orion
	Annotator
	50
	Kafon
	52
	5
	Phalanx
	54
	T
	56
	CyberZ
	ACTech
	58
	CAM
	Fensi
	60
	XT17/ADT
	ABEMAOTA
	62
	64
	6
	Engineering
	How
	We
	Work
	70
	70
	7b
	7
	XT17/ADT
	82
	*/
	for _, w := range resp.Blocks {
		if *w.BlockType == "WORD" {
			fmt.Println(*w.Text)
		}
	}

	// PDFのページ数を返す
	return 0, err
}


func main() {
	lambda.Start(s3Handler)
}
