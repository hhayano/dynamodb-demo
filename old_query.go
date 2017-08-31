package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	svc := dynamodb.New(session.New(aws.NewConfig().WithRegion("us-west-2")))
	queryInput := &dynamodb.QueryInput{

		KeyConditionExpression: aws.String("#NAME = :haji"),
		ProjectionExpression:   aws.String("Hometown"),
		ExpressionAttributeNames: map[string]*string{
			"#NAME": aws.String("Name"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":haji": {
				S: aws.String("Hajime Hayano"),
			},
		},

		TableName: aws.String("LifeDetails"),
	}
	result, _ := svc.Query(queryInput)
	fmt.Println(result)
}
