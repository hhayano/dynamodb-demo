package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	svc := dynamodb.New(session.New(aws.NewConfig().WithRegion("us-west-2")))
	updateInput := &dynamodb.UpdateItemInput{

		UpdateExpression: aws.String("SET Hometown = :city"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":city": {
				S: aws.String("Hilo"),
			},
		},

		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String("Hajime Hayano"),
			},
		},
		TableName:    aws.String("LifeDetails"),
		ReturnValues: aws.String("UPDATED_NEW"),
	}
	result, _ := svc.UpdateItem(updateInput)
	fmt.Println(result)
}
