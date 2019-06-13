package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func updateItem(id string, appid string, ts string) {
	fmt.Println("update", dynamotable)
	/*
		update := expression.Set(
			expression.Name("push"),
			expression.Value("push"),
		)

		expr, err := expression.NewBuilder().WithUpdate(update).Build()
		if err != nil {
			return err
		}
	*/

	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(dynamotable),
		Key: map[string]dynamodb.AttributeValue{

			"id": {
				S: aws.String(id),
			},
		},

		/*
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
		*/
		ExpressionAttributeNames: map[string]string{
			"#push": "push",
		},
		ExpressionAttributeValues: map[string]dynamodb.AttributeValue{
			":vals": {
				L: []dynamodb.AttributeValue{
					{
						M: map[string]dynamodb.AttributeValue{
							"appid": {
								S: aws.String(appid),
							},
							"ts": {
								S: aws.String(ts),
							},
						},
					},
				},
				//aws.String(ts),
			},
		},
		//UpdateExpression: aws.String("set #ts =  :ts_value"),
		UpdateExpression: aws.String("SET #push = list_append(#push, :vals)"),

		ReturnConsumedCapacity:      "NONE",
		ReturnItemCollectionMetrics: "NONE",
		ReturnValues:                "NONE",
	}
	// Example sending a request using PutItemRequest.
	req := dynamodbsvc.UpdateItemRequest(params)
	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
func putItem(id *string) {
	fmt.Println("putitem-", dynamotable)
	params := &dynamodb.PutItemInput{
		TableName: aws.String(dynamotable),
		Item: map[string]dynamodb.AttributeValue{

			"id": {
				S: id,
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
func getItem(id string) {
	fmt.Println("getitem-", id)
	params := &dynamodb.GetItemInput{
		TableName: aws.String(dynamotable),
		Key: map[string]dynamodb.AttributeValue{
			"id": {
				S: aws.String(string(id)),
			},
		},
		/*
			AttributesToGet: []string{
				"ami",
				"ts",
			},
		*/
		//	ConsistentRead:         aws.Bool(true),
		ReturnConsumedCapacity: "NONE",
	}
	// Examsple sending a request using PutItemRequest.
	req := dynamodbsvc.GetItemRequest(params)
	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Println("get", err)
	}
	fmt.Println("getItem:", resp)
	/*
		threadts = aws.StringValue(resp.Item["ts"].S)
		fmt.Println("threadts", threadts)
	*/
	for _, v := range resp.Item["push"].L {
		tmp := slackapps[aws.StringValue(v.M["appid"].S)]
		tmp.Ts = aws.StringValue(v.M["ts"].S)

		slackapps[aws.StringValue(v.M["appid"].S)] = tmp
	}

	//fmt.Println(slackapps)

	ami = aws.StringValue(resp.Item["ami"].S)
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
