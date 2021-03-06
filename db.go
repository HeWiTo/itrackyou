package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Declare a new DynamoDB instance. Note that this is safe for concurrent use.
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

func getItem(uniCode string) (*product, error) {
	// Prepare the input for the query.
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Product"),
		Key: map[string]*dynamodb.AttributeValue{
			"uni_code": {
				S: aws.String(uniCode),
			},
		},
	}

	// Retrieve the item from DynamoDB. If no matching item is found
	// return nil.
	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	prd := new(product)
	err = dynamodbattribute.UnmarshalMap(result.Item, prd)
	if err != nil {
		return nil, err
	}

	return prd, nil
}

func putItem(prd *product) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Product"),
		Item: map[string]*dynamodb.AttributeValue{
			"activate_date": {
				S: aws.String(prd.ActivateDate),
			},
			"brand": {
				S: aws.String(prd.Brand),
			},
			"detail": {
				B: []byte(prd.Detail),
			},
			"name": {
				S: aws.String(prd.Name),
			},
			"organization": {
				S: aws.String(prd.Organization),
			},
			"product_code": {
				S: aws.String(prd.ProductCode),
			},
			"product_date": {
				S: aws.String(prd.ProductDate),
			},
			"uni_code": {
				S: aws.String(prd.UniCode),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}
