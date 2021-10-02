package main

import (
	"regexp"
	"strconv"
	"strings"
)

type PageNumExtracter interface {
	extractPageNum([]string) (int, error)
}

type SimplePageNumExtracter struct {
}

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

func (s SimplePageNumExtracter) extractPageNum(detectWordList []string) (pageNum int, err error) {
	// 取得したワードのリストを一つの文字列にする
	detectWord := strings.Join(detectWordList, "")

	// 数字/数字 という条件に最初に合致する文字列を取得、最初の数字を現在見ているページ数とする
	r := regexp.MustCompile(`(\d)+\/\d+`)
	// FindStringSubmatchの戻り値は[最初の数字/数字 最初の数字]
	submatch := r.FindStringSubmatch(detectWord)
	// 最初の数字(ページ数)のみを取り出す
	pageNumString := submatch[1]

	// 取得したページ数をint型にキャストして返す
	pageNum, err = strconv.Atoi(pageNumString)
	return
}
