package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/your-org/your-lambda/internal/handler"
)

func main() {
	lambda.Start(handler.HandleAPIGatewayV2)
}
