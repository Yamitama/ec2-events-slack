package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func lambdaHandler(ctx context.Context) {

}

func main() {
	lambda.Start(lambdaHandler)
}
