package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Keys struct {
	PartitionKey string `json:"partition_key"`
	SortKey      string `json:"sort_key"`
}

func PutItem(tableName string, values interface{}) error {
	row, err := dynamodbattribute.MarshalMap(values)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      row,
		TableName: aws.String(tableName),
	}

	_, err = DynamoClient.PutItem(input)
	return err
}

func GetItem(tableName string, primaryKey string,primaryKeyValue string, secondaryKey string,secondaryKeyValue string,outputItem interface{}) (interface{}, error) {

	params := &dynamodb.GetItemInput{
		TableName:aws.String(tableName),
		Key:map[string]*dynamodb.AttributeValue{
			primaryKey: {
				S: aws.String(primaryKeyValue),
			},
			secondaryKey: {
				S: aws.String(secondaryKeyValue),
			},
		},
	}

	resp,err := DynamoClient.GetItem(params)

	if err != nil {
		return nil,err
	}
	err = dynamodbattribute.UnmarshalMap(resp.Item,outputItem)
	return outputItem,err
}
