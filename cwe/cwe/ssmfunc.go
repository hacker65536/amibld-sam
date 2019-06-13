package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func getParameter(sectype string) *ssm.GetParameterResponse {
	path := "/" + prefix + "/" + sectype + "/" + appID
	params := &ssm.GetParameterInput{
		Name:           aws.String(path),
		WithDecryption: aws.Bool(true),
	}

	req := ssmsvc.GetParameterRequest(params)
	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Println("error", err)
	}
	return resp

}

func getParametersByPath(sectype string) *ssm.GetParametersByPathResponse {

	path := "/" + prefix + "/" + sectype
	params := &ssm.GetParametersByPathInput{
		Path:           aws.String(path),
		Recursive:      aws.Bool(true),
		WithDecryption: aws.Bool(true),
	}
	req := ssmsvc.GetParametersByPathRequest(params)
	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	return resp

}
