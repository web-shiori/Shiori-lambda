package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PageNumExtracter interface {
	extractPageNum([]string) (int, error)
}

/*
	PDFのページ数を抽出する単純なロジック。
	数字/数字 という条件に最初に合致する文字列を取得、最初の数字を現在見ているページ数とする
*/
type SimplePageNumExtracter struct {
}

func (s SimplePageNumExtracter) extractPageNum(detectWordList []string) (pageNum int, err error) {
	// 取得したワードのリストを一つの文字列にする
	detectWord := strings.Join(detectWordList, "")

	r := regexp.MustCompile(`(\d+)\/\d+`)
	// FindStringSubmatchの戻り値は[最初の数字/数字 最初の数字]
	submatch := r.FindStringSubmatch(detectWord)
	// 最初の数字(ページ数)のみを取り出す
	pageNumString := "0"
	fmt.Println(submatch)
	if len(submatch) >= 2 {
		pageNumString = submatch[1]
	} else {
		fmt.Println("extracting page number was failed.")
	}

	pageNum, err = strconv.Atoi(pageNumString)
	return
}
