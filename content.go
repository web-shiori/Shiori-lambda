package main

import (
	"context"
	"fmt"
	"os"

	"github.com/web-shiori/web-snapshot-api/pkg/ws"
)

// PDFのページ数を保存する
func putPDFPageNum(contentID string, pageNum int) error {
	fmt.Println("------put pdf_page_num------")
	snapshotEmail := os.Getenv("SNAPSHOT_EMAIL")
	snapshotPassword := os.Getenv("SNAPSHOT_PASSWORD")
	fmt.Printf("Email: %s\n", snapshotEmail)
	fmt.Printf("Password: %s\n", snapshotPassword)

	c := ws.NewClient()
	auth := ws.AuthParams{
		EMail:    snapshotEmail,
		PassWord: snapshotPassword,
	}
	err := c.AuthService.SignIn(context.Background(), &auth)
	if err != nil {
		return err
	}

	content := ws.Content{
		ID:         contentID,
		PDFPageNum: pageNum,
	}
	_, err = c.ContentService.Update(context.Background(), &content)
	if err != nil {
		return err
	}
	fmt.Println("------put pdf_page_num was succeeded------")
	return nil
}
