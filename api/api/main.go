package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
)

var (
	prefix      = os.Getenv("PROJECTPREFIX")
	dynamotable = os.Getenv("DYNAMOTABLE")
	endpoint    = os.Getenv("DYNAMODB_ENDPOINT")
	signame     = os.Getenv("SSMSLACKSIGSECNAME")
	tokenname   = os.Getenv("SSMSLACKTOKENNAME")
	sig         []string
	// appID is slack's app id using to specify
	appID string
	token string
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("prefix", prefix)
	fmt.Println("dynamo", dynamotable)
	//fmt.Println("endpoint", endpoint)
	//authtoken := aws.StringValue(r.Parameter.Value)
	//fmt.Println(authtoken)

	ev, _ := json.Marshal(request)
	fmt.Println(string(ev))

	sigerr := false

	r := getParametersByPath(signame)
	//fmt.Println(r)

	// verify sig
	for _, v := range r.Parameters {
		s := aws.StringValue(v.Value)
		if er := VerifySig(s, request); er == nil {
			sigerr = true
			n := strings.Split(aws.StringValue(v.Name), "/")
			appID = n[3]
			r := getParameter(tokenname)
			token = aws.StringValue(r.Parameter.Value)
		}
	}
	/*
		// err test
		sig[0] = "620c1909213347efadf0365694760c6f"
		sig[1] = "c2c64d3eefadfefa164d86f45b57a488"
	*/

	if !sigerr {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("sig invalid")
	}

	contentTypeJSON := map[string]string{"Content-TYpe": "application/json"}
	//fmt.Printf("%s %s ", appID, token)

	//resbody := `{"aaa": "bbb"}`
	resbody := parseReqBody(request.Body)

	//	slackdialog(triggerID)

	//fmt.Println(listBuildsRequest())
	/*
		be := &Buildcfg{
			Account:  []string{""},
			Region:   []string{""},
			Platform: "amazonlinux2",
			UserName: "go",
			Commit:   "HEAD",
			Role:     []string{"httpd"},
			AMI:      "",
			DynamoDB: dynamotable,
		}
		br := startBuildRequest(be)
		fmt.Println(br)
		putItem(br.Build.Id)
	*/
	/*
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Hello, %v", string(ip)),
			StatusCode: 200,
		}, nil


	*/

	return events.APIGatewayProxyResponse{
		Body:       resbody,
		StatusCode: 200,
		Headers:    contentTypeJSON,
	}, nil
}

func main() {
	lambda.Start(handler)
}
