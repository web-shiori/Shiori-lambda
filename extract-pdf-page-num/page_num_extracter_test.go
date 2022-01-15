package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// extractPageNumのテスト
// 	* ページ数(数字/数字)が1つ含まれている文字列スライスが与えられた時、ページ数を返すこと
// 	* ページ数が含まれていない文字列スライスが与えられた時、0を返すこと
// 	* ページ数(数字/数字)が2つ以上含まれている文字列スライスが与えられた時、最初のページ数を返すこと
func TestExtractPageNum(t *testing.T) {
	tests := []struct {
		name           string
		detectWordList []string
		out            int
	}{
		{
			name:           `"6", "/", "48"`,
			detectWordList: []string{"6", "/", "48"},
			out:            6,
		},
		{
			name:           `"a", "/", "48"`,
			detectWordList: []string{"a", "/", "48"},
			out:            0,
		},
		{
			name:           `"6", "/", "48", "7", "/", "48"`,
			detectWordList: []string{"6", "/", "48", "7", "/", "48"},
			out:            6,
		},
		{
			name:           `"V100", "38", "/72", "-", "100%", "+", "=", "(Ma"`,
			detectWordList: []string{"V100", "38", "/72", "-", "100%", "+", "=", "(Ma"},
			out:            38,
		},
		{
			name:           `"CyberAgent", "Way", "2021", "(#####)", "103", "/", "149", "100%"`,
			detectWordList: []string{"CyberAgent", "Way", "2021", "(#####)", "103", "/", "149", "100%"},
			out:            103,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var simplePageNumExtractor SimplePageNumExtractor
			actual, err := simplePageNumExtractor.extractPageNum(tt.detectWordList)
			fmt.Printf("pdf page num: %d\n", actual)
			assert.Nil(t, err)
			assert.Equal(t, tt.out, actual)
		})
	}
}
