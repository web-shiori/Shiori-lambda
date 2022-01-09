package main

import (
	"fmt"
	"os"
)

// PDFのページ数を保存する
func putPDFPageNum(pageNum int) {
	//c := ws.NewClient()
	//auth := ws.AuthParams{
	//	EMail:    "",
	//	PassWord: "",
	//}
	//c.AuthService.SignIn(context.Background(), &auth)
	snapshotEmail := os.Getenv("SNAPSHOT_EMAIL")
	snapshotPassword := os.Getenv("SNAPSHOT_PASSWORD")
	fmt.Println(snapshotEmail)
	fmt.Println(snapshotPassword)
}
