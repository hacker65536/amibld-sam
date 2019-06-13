package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func putItem(id *string) {
	fmt.Println("putitem-", dynamotable)
	t := strconv.Itoa(int(time.Now().Unix()))
	params := &dynamodb.PutItemInput{
		TableName: aws.String(dynamotable),
		Item: map[string]dynamodb.AttributeValue{

			"id": {
				S: id,
			},
			"channel": {
				S: aws.String(slackMessage.Channel.ID),
			},
			"push": {
				L: []dynamodb.AttributeValue{},
			},
			"ttl": {
				N: aws.String(t),
			},
		},
	}
	// Example sending a request using PutItemRequest.
	req := dynamodbsvc.PutItemRequest(params)
	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
func listTables() {
	fmt.Println("table-", dynamotable)
	params := &dynamodb.ListTablesInput{}
	// Example sending a request using PutItemRequest.
	req := dynamodbsvc.ListTablesRequest(params)
	resp, err := req.Send(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
