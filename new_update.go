package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	expr "github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func main() {
	svc := dynamodb.New(session.New(aws.NewConfig().WithRegion("us-west-2")))

	update := expr.Set(expr.Name("Hometown"), expr.Value("Hilo"))
	expression, _ := expr.NewBuilder().WithUpdate(update).Build()

	updateInput := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String("Hajime Hayano"),
			},
		},
		ExpressionAttributeNames:  expression.Names(),
		ExpressionAttributeValues: expression.Values(),
		UpdateExpression:          expression.Update(),
		TableName:                 aws.String("LifeDetails"),
		ReturnValues:              aws.String("UPDATED_NEW"),
	}

	result, _ := svc.UpdateItem(updateInput)
	fmt.Println(result)
}
