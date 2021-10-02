package lambda_page_number_extract

import "testing"

type extractPageNumTest struct {
	detectWordList []string
	pageNum        int
}

// extractPageNumのテスト
// ページ数(数字/数字)が1つ含まれている文字列スライスが与えられた時、ページ数を返すこと
// ページ数が含まれていない文字列スライスが与えられた時、0を返すこと
// ページ数(数字/数字)が2つ以上含まれている文字列スライスが与えられた時、最初のページ数を返すこと
var extractPageNumTests = []extractPageNumTest{
	{[]string{"6", "/", "48"}, 6},
	{[]string{"a", "/", "48"}, 0},
	{[]string{"6", "/", "48", "7", "/", "48"}, 6},
}

func TestExtractPageNum(t *testing.T) {
	var simplePageNumExtracter SimplePageNumExtracter
	for i, test := range extractPageNumTests {
		actual, err := simplePageNumExtracter.extractPageNum(test.detectWordList)
		if err != nil {
			t.Errorf("#%d: error occured.\nMSG:\n\t%s", i, err)
		}
		if actual != test.pageNum {
			t.Errorf("#%d: got: %d want: %d", i, actual, test.pageNum)
		}
	}
}
