package main

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
)

var (
	slackapps = map[string]slackApp{}
)

type slackApp struct {
	AppID      string
	Token      string
	Webhookurl string
	Sigsec     string
	ChannelID  string
	Ts         string
}

func initSlackApp() {

	r := getParametersByPath(webhookname)
	for _, v := range r.Parameters {
		n := strings.Split(aws.StringValue(v.Name), "/")
		appID := n[3]
		app := slackApp{}
		app.AppID = appID
		app.Webhookurl = aws.StringValue(v.Value)
		slackapps[appID] = app

	}

	r = getParametersByPath(tokenname)
	for _, v := range r.Parameters {
		n := strings.Split(aws.StringValue(v.Name), "/")
		appID := n[3]
		tmp := slackapps[appID]
		tmp.Token = aws.StringValue(v.Value)
		slackapps[appID] = tmp
	}

	r = getParametersByPath(channelname)
	for _, v := range r.Parameters {
		n := strings.Split(aws.StringValue(v.Name), "/")
		appID := n[3]
		tmp := slackapps[appID]
		tmp.ChannelID = aws.StringValue(v.Value)
		slackapps[appID] = tmp
	}
}

/*
func getSec() {

	r := getParameter(tokenname)

	tokens = append(tokens, aws.StringValue(r.Parameter.Value))
	//fmt.Println(tokens)

	r = getParameter(sigsecname)
	//fmt.Println(r)

	sigsecs = append(sigsecs, aws.StringValue(r.Parameter.Value))
	//fmt.Println(sigsecs)
}
*/
