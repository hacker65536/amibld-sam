package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

var (
	api = slack.New(token)
)

func getuser() {
	user, err := api.GetUserInfo("xxxxx")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}

func slackdialog(triggerID string) {
	api = slack.New(token)
	dialog := slack.Dialog{
		Title:      "AMIBLD",
		CallbackID: "amiorder",
		Elements: []slack.DialogElement{
			slack.DialogInputSelect{
				Value: "amazonlinux2",
				DialogInput: slack.DialogInput{
					Label:       "Platform",
					Type:        slack.InputTypeSelect,
					Name:        "platform",
					Placeholder: "Select a platform",
				},
				Options: []slack.DialogSelectOption{
					{
						Label: "amazonlinux",
						Value: "amazonlinux",
					},
					{
						Label: "amazonlinux2",
						Value: "amazonlinux2",
					},
					{
						Label: "centos6",
						Value: "centos6",
					},
					{
						Label: "centos7",
						Value: "centos7",
					},
					{
						Label: "ubuntu18",
						Value: "ubuntu18",
					},
				},
			},
			slack.TextInputElement{
				DialogInput: slack.DialogInput{
					Label:       "Roles",
					Type:        slack.InputTypeTextArea,
					Name:        "roles",
					Placeholder: "base,php",
					Optional:    true,
				},
				Value: "base,php",
			},
			slack.TextInputElement{
				DialogInput: slack.DialogInput{
					Label:       "CommitID(revsion)",
					Type:        slack.InputTypeText,
					Name:        "commit",
					Optional:    true,
					Placeholder: "HEAD",
				},
			},
			slack.TextInputElement{
				DialogInput: slack.DialogInput{
					Label:       "Share account",
					Type:        slack.InputTypeText,
					Name:        "accounts",
					Optional:    true,
					Placeholder: "00000000000",
				},
			},
			slack.TextInputElement{
				DialogInput: slack.DialogInput{
					Label:       "Copy to another region",
					Type:        slack.InputTypeText,
					Name:        "regions",
					Optional:    true,
					Placeholder: "us-east-2",
				},
			},
			slack.TextInputElement{
				DialogInput: slack.DialogInput{
					Label:       "sourceAMI",
					Type:        slack.InputTypeText,
					Name:        "ami",
					Optional:    true,
					Placeholder: "ami-0000000000",
				},
			},
		},
	}
	err := api.OpenDialog(triggerID, dialog)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

/*
func helpSlackBlock() string {
	divSection := slack.NewDividerBlock()
	headerText := slack.NewTextBlockObject("mrkdwn", "*Usage:*\n/command os=amazonlinux account=000000000000 region=us-west-2 role=httpd", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	fieldSlice := make([]*slack.TextBlockObject, 0)

	fieldSlice = append(fieldSlice, slack.NewTextBlockObject(
		"mrkdwn",
		"*os:(platform)*\nos=amazonlinux2",
		false,
		false))

	fieldSlice = append(fieldSlice, slack.NewTextBlockObject(
		"mrkdwn",
		"*role:(ansible role)*\nrole=base,httpd",
		false,
		false))

	fieldSlice = append(fieldSlice, slack.NewTextBlockObject(
		"mrkdwn",
		"*region:(copy to specific region)*\nregion=us-east-2 ",
		false,
		false))

	fieldSlice = append(fieldSlice, slack.NewTextBlockObject(
		"mrkdwn",
		"*account:(share to another account)*\naccount=012345678901 ",
		false,
		false))

	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)
	ExSection := slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "\n*available os* amazonlinux / amazonlinux2 / centos6 / centos7 / ubuntu18", false, false), nil, nil)
	msg := slack.NewBlockMessage(
		divSection,
		headerSection,
		fieldsSection,
		ExSection,
	)

	b, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		fmt.Println(err)
		//return
	}

	fmt.Println("help", string(b))
	return string(b)
}
*/
/*
func confirmSlackBlock(be *Buildcfg) string {

	t := time.Now().Unix()
	// Header Section
	divSection := slack.NewDividerBlock()
	headerText := slack.NewTextBlockObject("mrkdwn", "*Confirm your request* \n <!date^"+strconv.FormatInt(t, 10)+"^Posted {date_num} {time_secs}| 2014-02-18 6:39:42 AM PST>", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Fields
	platformField := slack.NewTextBlockObject("mrkdwn", "*Platform:*\n"+be.Platform, false, false)
	userNameField := slack.NewTextBlockObject("mrkdwn", "*User:*\n"+be.UserName, false, false)
	shareAccountField := slack.NewTextBlockObject("mrkdwn", "*ShareAccount:*\n"+be.ShareAccount, false, false)
	copyRegionField := slack.NewTextBlockObject("mrkdwn", "*CopyRegion:*\n"+be.CopyRegion, false, false)
	commitIDField := slack.NewTextBlockObject("mrkdwn", "*CommitID:*\n"+be.CommitID, false, false)
	rolesField := slack.NewTextBlockObject("mrkdwn", "*Roles:*\n"+fmt.Sprintf("%s", be.Roles), false, false)

	fieldSlice := make([]*slack.TextBlockObject, 0)

	fieldSlice = append(fieldSlice, platformField)
	fieldSlice = append(fieldSlice, userNameField)
	fieldSlice = append(fieldSlice, shareAccountField)
	fieldSlice = append(fieldSlice, copyRegionField)
	fieldSlice = append(fieldSlice, commitIDField)
	fieldSlice = append(fieldSlice, rolesField)

	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

	// Approve and Deny Buttons
		approveBtnTxt := slack.NewTextBlockObject("plain_text", "Build", true, false)
		approveBtn := slack.NewButtonBlockElement("build", "build", approveBtnTxt)
		approveBtn.Style = "primary"
		approveBtn.Confirm = slack.NewConfirmationBlockObject(
			slack.NewTextBlockObject("plain_text", "Are you sure?", false, false),
			slack.NewTextBlockObject("mrkdwn", "Wouldn't you prefer a good game of _chess_?", false, false),
			slack.NewTextBlockObject("plain_text", "Do it", false, false),
			slack.NewTextBlockObject("plain_text", "stop", false, false),
		)

		denyBtnTxt := slack.NewTextBlockObject("plain_text", "Cancel", true, false)
		denyBtn := slack.NewButtonBlockElement("", "cancel", denyBtnTxt)
		denyBtn.Style = "danger"

		actionBlock := slack.NewActionBlock("", approveBtn, denyBtn)

	// Build Message with blocks created above

	msg := slack.NewBlockMessage(
		divSection,
		headerSection,
		fieldsSection,
	//	actionBlock,
	)

	b, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		fmt.Println(err)
		//return
	}

	fmt.Println("conf", string(b))
	return string(b)

}
*/
