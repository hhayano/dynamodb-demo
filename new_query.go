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

	keyCond := expr.Key("Name").Equal(expr.Value("Hajime Hayano"))
	proj := expr.NamesList(expr.Name("Hometown"))
	expression, _ := expr.NewBuilder().WithKeyCondition(keyCond).WithProjection(proj).Build()

	queryInput := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expression.Names(),
		ExpressionAttributeValues: expression.Values(),
		KeyConditionExpression:    expression.KeyCondition(),
		ProjectionExpression:      expression.Projection(),
		TableName:                 aws.String("LifeDetails"),
	}

	result, _ := svc.Query(queryInput)

	fmt.Println(result)
}
