package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	prefix     = os.Getenv("PROJECTPREFIX")
	tokenname  = os.Getenv("SSMSLACKTOKENNAME")
	sigsecname = os.Getenv("SSMSLACKSIGSECNAME")
	appID      = "AHQQ4HCMP"
	sigsecs    []string
	tokens     []string
)

func handler(event events.CloudWatchEvent) (string, error) {
	j, _ := json.Marshal(event)
	fmt.Println(string(j))

	ed := event.Detail
	fmt.Println("event", string(ed))

	getSec()

	var be BuildEvent
	_ = json.Unmarshal((json.RawMessage)(ed), &be)
	fmt.Println("be", be)

	att := genAttachment(be)
	postMes(tokens[0], att)

	return fmt.Sprintf("event: %v!", "aaa"), nil

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
