package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var DynamoClient *dynamodb.DynamoDB


func NewClient() {

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("AKIA6H37YGCAUJ4VRVJB", "Ck5PyAtaT4NY7BTz3kXF2JQ3nBIS+kfkJbvCX2qz", ""),
	}))

	/*sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:aws.String("http://localhost:8000"),
		Region:aws.String("us-east-1"),
	}))*/

	DynamoClient = dynamodb.New(sess)
}

