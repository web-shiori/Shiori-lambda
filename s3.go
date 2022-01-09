package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"strings"
)

// S3に保存されたオブジェクトからコンテンツIDを取得する。
// S3では"/pdf/コンテンツID/〇〇.png"というフォルダ構成になっている。
// ↑からコンテンツIDを取得する
func getContentID(object events.S3Object) (string, error) {
	ss := strings.Split(object.Key, "/")
	if len(ss) != 3 {
		return "", fmt.Errorf("%s: Invalid object", ss)
	}
	return ss[1], nil
}

// S3からスクリーンショットを取得

// S3からスクリーンショットを削除
