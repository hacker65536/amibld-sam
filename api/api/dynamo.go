package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	dynamodbsvc *dynamodb.Client
)

func init() {
	/*
		ddbCustResolverFn := func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           endpoint,
				SigningRegion: "us-east-1",
			}, nil
		}

	*/
	cfgCp := cfg.Copy()
	fmt.Println(endpoint)
	if endpoint != "" {
		cfgCp.EndpointResolver = aws.ResolveWithEndpointURL(endpoint)
		fmt.Println("end", endpoint)
	} else {

	}
	dynamodbsvc = dynamodb.New(cfgCp)
}
