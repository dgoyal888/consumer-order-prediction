package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var DynamoClient *dynamodb.DynamoDB


func NewClient() {

	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:aws.String("http://localhost:8000"),
		Region:aws.String("us-east-1"),
	}))

	DynamoClient = dynamodb.New(sess)
}

