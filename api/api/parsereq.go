package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/nlopes/slack"
)

var (
	triggerID    string
	slackMessage slack.InteractionCallback
)

type DialogError struct {
	Serrs []Serr `json:"errors"`
}

type Serr struct {
	Name  string `json:"name"`
	Error string `json:"error"`
}

func parseReqBody(body string) string {

	//hash := make(map[string]string)
	//fmt.Println("ispayload", string(body)[:8])

	// if start with payload(dialog request)
	if string(body)[:8] == "payload=" {
		jsonStr, _ := url.QueryUnescape(string(body)[8:])
		if err := json.Unmarshal([]byte(jsonStr), &slackMessage); err != nil {
			fmt.Println(err)
		}

		//j, _ := json.Marshal(message)
		//fmt.Println(string(j))

		r := startBuildFromSlack()
		fmt.Println("r", r)
		return r

	}

	if string(body)[:6] == "token=" {
		d, e := url.ParseQuery(body)
		if e != nil {
			log.Fatal(e)
		}

		//fmt.Println(d["trigger_id"])
		triggerID = d["trigger_id"][0]
		slackdialog(triggerID)
	}

	return `{"text": "ok"}`
}

func startBuildFromSlack() string {
	fmt.Println("user:", slackMessage.User.Name)
	fmt.Println("platform:", slackMessage.Submission["platform"])
	fmt.Println("regions:", slackMessage.Submission["regions"])
	fmt.Println("roles:", slackMessage.Submission["roles"])
	fmt.Println("accounts:", slackMessage.Submission["accounts"])
	fmt.Println("commit:", slackMessage.Submission["commit"])
	fmt.Println("ami:", slackMessage.Submission["ami"])
	fmt.Println("channel:", slackMessage.Channel.ID)

	roles, regions, accounts := []string{}, []string{}, []string{}

	roles = strings.Split(slackMessage.Submission["roles"], ",")
	if slackMessage.Submission["roles"] == "" {
		roles = []string{"base"}
	}

	regions = strings.Split(slackMessage.Submission["regions"], ",")
	if slackMessage.Submission["regions"] == "" {
		regions = []string{}
	}
	accounts = strings.Split(slackMessage.Submission["accounts"], ",")
	if slackMessage.Submission["accounts"] == "" {
		accounts = []string{}
	}

	d := &DialogError{}
	if slackMessage.Submission["account"] == "0000" {
		d.Serrs = append(d.Serrs, Serr{Name: "account", Error: "invalid account"})
	}

	if len(d.Serrs) != 0 {
		j, _ := json.Marshal(d)
		return string(j)
	}
	be := &Buildcfg{
		Account:     accounts,
		Region:      regions,
		Platform:    slackMessage.Submission["platform"],
		UserName:    slackMessage.User.Name,
		Commit:      slackMessage.Submission["commit"],
		Role:        roles,
		AMI:         slackMessage.Submission["ami"],
		DynamoTable: dynamotable,
	}
	br := startBuildRequest(be)
	fmt.Println(br)
	putItem(br.Build.Id)
	return "{}"
}
