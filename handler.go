package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func excuteFunction()  {
	fmt.Println("Hello, Lambda")
}

func main() {
	lambda.Start(excuteFunction)
}
