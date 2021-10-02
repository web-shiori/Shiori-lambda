package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func lambdaHandler() {
	fmt.Println("TODO: ページ数を保存する")
}

func main() {
	lambda.Start(lambdaHandler)
}
