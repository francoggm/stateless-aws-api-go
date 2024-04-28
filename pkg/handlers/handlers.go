package handlers

import (
	"aws-stateless/pkg/user"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

var ErrorMethodNotAllowed = "method not allowed"

func GetUser(req events.APIGatewayProxyRequest, db *dynamodb.Client, tableName string) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(email, db, tableName)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{
				ErrorMsg: aws.String(err.Error()),
			})
		}

		if result == nil {
			return apiResponse(http.StatusNotFound, nil)
		}

		return apiResponse(http.StatusOK, result)
	}

	result, err := user.FetchUsers(db, tableName)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayProxyRequest, db *dynamodb.Client, tableName string) (*events.APIGatewayProxyResponse, error) {
	result, err := user.CreateUser(req, db, tableName)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	return apiResponse(http.StatusCreated, result)
}

func UpdateUser(req events.APIGatewayProxyRequest, db *dynamodb.Client, tableName string) (*events.APIGatewayProxyResponse, error) {
	result, err := user.UpdateUser(req, db, tableName)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	return apiResponse(http.StatusCreated, result)
}

func DeleteUser(req events.APIGatewayProxyRequest, db *dynamodb.Client, tableName string) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]
	if len(email) <= 0 {
		return apiResponse(http.StatusNotFound, nil)
	}

	err := user.DeleteUser(email, db, tableName)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	return apiResponse(http.StatusOK, nil)
}

func UnhandleMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
