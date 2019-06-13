package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	prefix      = os.Getenv("PROJECTPREFIX")
	tokenname   = os.Getenv("SSMSLACKTOKENNAME")
	webhookname = os.Getenv("SSMSLACKWEBHOOKNAME")
	channelname = os.Getenv("SSMSLACKCHANNELNAME")
	dynamotable = os.Getenv("DYNAMOTABLE")
	endpoint    = os.Getenv("DYNAMODB_ENDPOINT")
	//sigsecname  = os.Getenv("SSMSLACKSIGSECNAME")
	//sigsecs     []string
	//tokens      []string
	//webhooks    []string
	threadts string
	//appID       string
	ami string
)

func handler(event events.CloudWatchEvent) (string, error) {
	fmt.Println("dynamo", dynamotable)
	fmt.Println("endpoint", endpoint)
	j, _ := json.Marshal(event)
	fmt.Println(string(j))

	ed := event.Detail
	fmt.Println("event", string(ed))

	// init token sig
	initSlackApp()

	var be BuildEvent
	_ = json.Unmarshal((json.RawMessage)(ed), &be)

	//fmt.Println("be", be)
	//fmt.Println("sourceversion", be.AdditionalInformation.SourceVersion)

	id := strings.Split(be.BuildID, "/")
	getItem(id[1])

	att := genAttachment(be)

	var ts string
	for _, v := range slackapps {
		threadts = v.Ts
		ts = postMes(v.Token, v.ChannelID, att)
		updateItem(id[1], v.AppID, ts)
	}

	return fmt.Sprintf("event: %v!", ts), nil

}

func main() {
	lambda.Start(handler)
}

func searchValue(name string, env []EnvironmentVariables) string {
	//fmt.Println(env)
	sort.Slice(env, func(i, j int) bool {
		return env[i].Name <= env[j].Name
	})

	n := name
	idx := sort.Search(len(env), func(i int) bool {
		return string(env[i].Name) >= n
	})

	if env[idx].Name == n {
		return env[idx].Value
	}
	return ""
}
