package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

func lambdaHandler(ctx context.Context, event events.S3Event) {
	// S3から画像オブジェクト取得
	fmt.Println("S3から画像オブジェクト取得")
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	inputBucket := event.Records[0].S3.Bucket.Name
	inputKey := event.Records[0].S3.Object.Key
	fmt.Println(inputBucket)
	fmt.Println(inputKey)

	// Textract API呼び出し
	fmt.Println("Textract API呼び出し")
	textractSession := textract.New(sess)
	input := &textract.DetectDocumentTextInput{
		Document: &textract.Document{
			S3Object: &textract.S3Object{
				Bucket: aws.String(inputBucket),
				Name:   aws.String(inputKey),
			},
		},
	}
	resp, err := textractSession.DetectDocumentText(input)
	if err != nil {
		panic(err)
	}

	// レスポンスからテキストデータを抽出
	fmt.Println("レスポンスからテキストデータを抽出")
	var output string
	for _, w := range resp.Blocks {
		if *w.BlockType == "LINE" {
			output += *w.Text
			output += "\n"
		}
	}
	fmt.Println(output)

	// S3に書き込み
	fmt.Println("S3に書き込み")
	outputBucket := "" // 出力用バケット名
	outputObjectKey := inputKey + time.Now().String() + ".txt"
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(outputBucket),
		Key:    aws.String(outputObjectKey),
		Body:   strings.NewReader(output),
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	lambda.Start(lambdaHandler)
}
