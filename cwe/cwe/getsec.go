package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

func getSecs() {

	r := getParametersByPath(tokenname)
	for _, v := range r.Parameters {
		tokens = append(tokens, aws.StringValue(v.Value))
	}
	//fmt.Println(tokens)

	r = getParametersByPath(sigsecname)
	for _, v := range r.Parameters {
		sigsecs = append(sigsecs, aws.StringValue(v.Value))
	}
	//fmt.Println(sigsecs)
}

func getSec() {

	r := getParameter(tokenname)

	tokens = append(tokens, aws.StringValue(r.Parameter.Value))
	//fmt.Println(tokens)

	r = getParameter(sigsecname)
	//fmt.Println(r)

	sigsecs = append(sigsecs, aws.StringValue(r.Parameter.Value))
	//fmt.Println(sigsecs)
}
