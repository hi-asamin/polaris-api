package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"polaris-api/infrastructure"
	"polaris-api/infrastructure/router"
)

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	infrastructure.NewDatabase()

	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		// Lambda環境
		r := router.CreateRouter()
		ginLambda = ginadapter.New(r)
		lambda.Start(Handler)
	} else {
		// ローカル環境
		router.Init()
	}
}
