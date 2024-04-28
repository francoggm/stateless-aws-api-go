package user

import (
	"aws-stateless/pkg/validators"
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var (
	ErrorFailedFetch        = "failed to fetch record"
	ErrorFailedUnmarshal    = "failed to unmarshal record"
	ErrorFailedFetchAll     = "failed to fetch all table"
	ErrorFailedUnmarshalAll = "failed to unmarshal all records"
	ErrorFailedMarshall     = "failed to marshall record"
	ErrorFailedPutItem      = "failed to put item"
	ErrorFailedUpdateItem   = "failed to update item"
	ErrorFailedDeleteItem   = "failed to delete item"
	ErrorInvalidUserData    = "invalid user data body"
	ErrorInvalidUserEmail   = "invalid user email"
	ErrorUserAlreadyExists  = "user already exists"
	ErrorUserDontExists     = "user don't exists"
)

func FetchUser(email string, db *dynamodb.Client, tableName string) (*User, error) {
	keys, err := attributevalue.MarshalMap(map[string]any{
		"email": email,
	})
	if err != nil {
		return nil, errors.New(ErrorFailedMarshall)
	}

	result, err := db.GetItem(context.Background(), &dynamodb.GetItemInput{
		Key:       keys,
		TableName: &tableName,
	})
	if err != nil {
		return nil, errors.New(ErrorFailedFetch)
	}

	user := &User{}
	err = attributevalue.UnmarshalMap(result.Item, user)
	if err != nil {
		return nil, errors.New(ErrorFailedUnmarshal)
	}

	return user, nil
}

func FetchUsers(db *dynamodb.Client, tableName string) ([]*User, error) {
	result, err := db.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return nil, errors.New(ErrorFailedFetchAll)
	}

	users := make([]*User, 0)

	err = attributevalue.UnmarshalListOfMaps(result.Items, users)
	if err != nil {
		return nil, errors.New(ErrorFailedUnmarshalAll)
	}

	return users, nil
}

func CreateUser(req events.APIGatewayProxyRequest, db *dynamodb.Client, tableName string) (*User, error) {
	var u User

	if err := json.Unmarshal([]byte(req.Body), &u); err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}

	if !validators.IsEmailValid(u.Email) {
		return nil, errors.New(ErrorInvalidUserEmail)
	}

	currentUser, _ := FetchUser(u.Email, db, tableName)
	if currentUser != nil {
		return nil, errors.New(ErrorUserAlreadyExists)
	}

	av, err := attributevalue.MarshalMap(u)
	if err != nil {
		return nil, errors.New(ErrorFailedMarshall)
	}

	_, err = db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	})
	if err != nil {
		return nil, errors.New(ErrorFailedPutItem)
	}

	return &u, nil
}

func UpdateUser(req events.APIGatewayProxyRequest, db *dynamodb.Client, tableName string) (*User, error) {
	var u User

	if err := json.Unmarshal([]byte(req.Body), &u); err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}

	if currentUser, _ := FetchUser(u.Email, db, tableName); currentUser == nil {
		return nil, errors.New(ErrorUserDontExists)
	}

	av, err := attributevalue.MarshalMap(u)
	if err != nil {
		return nil, errors.New(ErrorFailedMarshall)
	}

	_, err = db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})
	if err != nil {
		return nil, errors.New(ErrorFailedUpdateItem)
	}

	return &u, nil
}

func DeleteUser(email string, db *dynamodb.Client, tableName string) error {
	keys, err := attributevalue.MarshalMap(map[string]any{
		"email": email,
	})
	if err != nil {
		return errors.New(ErrorFailedMarshall)
	}

	_, err = db.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       keys,
	})
	if err != nil {
		return errors.New(ErrorFailedDeleteItem)
	}

	return nil
}
