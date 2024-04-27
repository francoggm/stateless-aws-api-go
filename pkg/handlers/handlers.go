package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetUser(req events.APIGatewayProxyRequest, db *dynamodb.Client, tableName string) (*events.APIGatewayProxyResponse, error) {
	return nil, nil
}

func CreateUser(req events.APIGatewayProxyRequest, db *dynamodb.Client, tableName string) (*events.APIGatewayProxyResponse, error) {
	return nil, nil
}

func UpdateUser(req events.APIGatewayProxyRequest, db *dynamodb.Client, tableName string) (*events.APIGatewayProxyResponse, error) {
	return nil, nil
}

func DeleteUser(req events.APIGatewayProxyRequest, db *dynamodb.Client, tableName string) (*events.APIGatewayProxyResponse, error) {
	return nil, nil
}

func UnhandleMethod(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return nil, nil
}
