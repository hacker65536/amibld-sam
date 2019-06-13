package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/nlopes/slack"
)

func genAttachment(be BuildEvent) slack.Attachment {
	user := searchValue("USER_NAME", be.AdditionalInformation.Environment.EnvironmentVariables)
	platform := searchValue("PLATFORM", be.AdditionalInformation.Environment.EnvironmentVariables)
	layout := "Jan 2, 2006 15:04:05 PM"

	attachment := slack.Attachment{
		Color:      "good",
		Fallback:   "You successfully posted by Incoming Webhook URL!",
		AuthorName: "Author: " + user,
		//AuthorSubname: "github.com",
		AuthorLink: "https://github.com/nlopes/slack",
		Title:      "",
		TitleLink:  be.AdditionalInformation.Logs.DeepLink,
		Footer:     "slack api",
		FooterIcon: "https://platform.slack-edge.com/img/default_application_icon.png",
		Ts:         json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
		Fields: []slack.AttachmentField{
			{
				Title: "BuildID",
				Value: be.BuildID,
				Short: false,
			},
			{
				Title: "Platform",
				Value: platform,
				Short: false,
			},
		},
	}

	switch be.BuildStatus {
	case "SUCCEEDED":
		fmt.Println("-success")
		attachment.Color = "#00BFFF"
		attachment.Footer = "build-start-time"
		attachment.Title = "build-finished"
		t, _ := time.Parse(layout, be.AdditionalInformation.BuildStartTime)
		attachment.Ts = json.Number(strconv.FormatInt(t.Unix(), 10))
	case "STOPPED":
		fmt.Println("stopped")
		fmt.Println(be)
	case "FAILED":
		fmt.Println("-failed")
		attachment.Color = "danger"
		attachment.Footer = "build-start-time"
		attachment.Title = "build-failed"
		t, _ := time.Parse(layout, be.AdditionalInformation.BuildStartTime)
		attachment.Ts = json.Number(strconv.FormatInt(t.Unix(), 10))
	case "IN_PROGRESS":
		fmt.Println("-inprogress")
		attachment.Color = "#CEECF5"
		attachment.Footer = "build-start-time"
		attachment.Title = "build-start"
		t, _ := time.Parse(layout, be.AdditionalInformation.BuildStartTime)
		attachment.Ts = json.Number(strconv.FormatInt(t.Unix(), 10))
	}

	return attachment
}

func postMes(token string, attachment slack.Attachment) {
	api := slack.New(token)

	channelID, timestamp, err := api.PostMessage("CJ7QZ2631", slack.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

/*
func slackPostWebhook(be BuildEvent) {

	user := searchValue("USER_NAME", be.AdditionalInformation.Environment.EnvironmentVariables)
	platform := searchValue("PLATFORM", be.AdditionalInformation.Environment.EnvironmentVariables)
	//rev := searchValue("PLATFORM", be.AdditionalInformation.Environment.EnvironmentVariables)

	fmt.Println("deep-", be.AdditionalInformation.Logs.DeepLink)
	layout := "Jan 2, 2006 15:04:05 PM"
	attachment := slack.Attachment{
		Color:      "good",
		Fallback:   "You successfully posted by Incoming Webhook URL!",
		AuthorName: "Author: " + user,
		//AuthorSubname: "github.com",
		AuthorLink: "https://github.com/nlopes/slack",
		Title:      "",
		TitleLink:  be.AdditionalInformation.Logs.DeepLink,
		//		Text:       "aaa:",
		Footer:     "slack api",
		FooterIcon: "https://platform.slack-edge.com/img/default_application_icon.png",
		Ts:         json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
		Fields: []slack.AttachmentField{
			{
				Title: "BuildID",
				Value: be.BuildID,
				Short: false,
			},
			{
				Title: "Platform",
				Value: platform,
				Short: false,
			},
		},
	}

	//t, _ := json.Marshal(be)
	switch be.BuildStatus {
	case "SUCCEEDED":
		fmt.Println("-success")
		attachment.Color = "#00BFFF"
		attachment.Footer = "build-start-time"
		attachment.Title = "build-finished"
		t, _ := time.Parse(layout, be.AdditionalInformation.BuildStartTime)
		attachment.Ts = json.Number(strconv.FormatInt(t.Unix(), 10))
	case "STOPPED":
		fmt.Println("stopped")
		fmt.Println(be)
	case "FAILED":
		fmt.Println("-failed")
		attachment.Color = "danger"
		attachment.Footer = "build-start-time"
		attachment.Title = "build-failed"
		t, _ := time.Parse(layout, be.AdditionalInformation.BuildStartTime)
		attachment.Ts = json.Number(strconv.FormatInt(t.Unix(), 10))
	case "IN_PROGRESS":
		fmt.Println("-inprogress")
		attachment.Color = "#CEECF5"
		attachment.Footer = "build-start-time"
		attachment.Title = "build-start"
		t, _ := time.Parse(layout, be.AdditionalInformation.BuildStartTime)
		attachment.Ts = json.Number(strconv.FormatInt(t.Unix(), 10))
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.PostWebhook(WebhookURL, &msg)
	if err != nil {
		fmt.Println(err)
	}

}
*/
