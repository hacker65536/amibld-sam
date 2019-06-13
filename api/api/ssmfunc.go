package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func getParameter(sectype string) *ssm.GetParameterResponse {

	name := "/" + prefix + "/" + sectype + "/" + appID
	input := &ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	}

	//	fmt.Println("ssminput", input)
	req := ssmsvc.GetParameterRequest(input)
	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Println("error", err)
	}
	//fmt.Println(resp)
	//fmt.Println("token:", aws.StringValue(resp.Parameter.Value))
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
