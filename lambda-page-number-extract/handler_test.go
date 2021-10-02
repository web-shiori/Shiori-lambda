package main

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/textract"
)

// s3Handlerのテスト
// Q: 何をテストすれば良い？

// detectPageNumberのテスト
// bucket名とオブジェクトキーが与えられた時、ページ数を返すこと
// TODO: textractのモックの作り方を勉強する

// detectDocumentTextOutputToStringSliceのテスト
// textract.DetectDocumentTextOutputが与えられた時、それを文字列スライスにして返すこと
func TestDetectDocumentTextOutputToStringSlice(t *testing.T) {
	detectDocumentTextOutput := textract.DetectDocumentTextOutput{
		Blocks: []*textract.Block{
			{
				BlockType:       aws.String("WORD"),
				ColumnIndex:     nil,
				ColumnSpan:      nil,
				Confidence:      nil,
				EntityTypes:     nil,
				Geometry:        nil,
				Id:              nil,
				Page:            nil,
				Relationships:   nil,
				RowIndex:        nil,
				RowSpan:         nil,
				SelectionStatus: nil,
				Text:            aws.String("word"),
				TextType:        aws.String("PRINTED"),
			},
		},
		DetectDocumentTextModelVersion: aws.String("1.0"),
		DocumentMetadata: &textract.DocumentMetadata{
			Pages: aws.Int64(1),
		},
	}
	expected := []string{"word"}

	actual, err := detectDocumentTextOutputToStringSlice(&detectDocumentTextOutput)
	if err != nil {
		t.Errorf("error occured.\nMSG:\n\t%s", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got: %v want: %v", actual, expected)
	}
}
