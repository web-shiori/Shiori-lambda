package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/service/textract"

	"github.com/aws/aws-lambda-go/events"
)

type S3Service interface {
	getTextractS3Object() *textract.S3Object
	deleteObject() error
	getContentID() (string, error)
}

type s3Service struct {
	record events.S3EventRecord
}

var _ S3Service = (*s3Service)(nil)

// S3からスクリーンショットを取得
func (s *s3Service) getTextractS3Object() *textract.S3Object {
	panic("implement me")
	return nil
}

// S3からスクリーンショットを削除
func (s *s3Service) deleteObject() error {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &s.record.S3.Bucket.Name,
		Key:    &s.record.S3.Object.Key,
	})
	return err
}

// S3に保存されたオブジェクトからコンテンツIDを取得する。
// S3では"/pdf/コンテンツID/〇〇.png"というフォルダ構成になっている。
// ↑からコンテンツIDを取得する
func (s *s3Service) getContentID() (string, error) {
	ss := strings.Split(s.record.S3.Object.Key, "/")
	if len(ss) != 3 {
		return "", fmt.Errorf("%s: Invalid object", ss)
	}
	return ss[1], nil
}
