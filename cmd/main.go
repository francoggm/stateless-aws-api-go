package main

import (
	"aws-stateless/pkg/handlers"
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const (
	tableName = "serverless-users"
)

var svc *dynamodb.Client

func main() {
	region := "us-east-1" //os.Getenv("AWS_REGION")
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		log.Panic(err)
	}

	svc = dynamodb.NewFromConfig(cfg)
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetUser(req, svc, tableName)
	case "POST":
		return handlers.CreateUser(req, svc, tableName)
	case "PUT":
		return handlers.UpdateUser(req, svc, tableName)
	case "DELETE":
		return handlers.DeleteUser(req, svc, tableName)
	default:
		return handlers.UnhandleMethod()
	}
}
