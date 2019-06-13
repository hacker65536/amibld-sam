package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nlopes/slack"
)

// VerifySig is verification of slack request
func VerifySig(sig string, request events.APIGatewayProxyRequest) error {

	httpHd := make(map[string][]string)
	for k, v := range request.Headers {
		//fmt.Println(k)
		//fmt.Println(v)
		httpHd[k] = []string{v}
	}
	//fmt.Println(httpHd)
	/*
		httpHd = http.Header{
			"X-Slack-Signature":         {request.Headers["X-Slack-Signature"]},
			"X-Slack-Request-Timestamp": {request.Headers["X-Slack-Request-Timestamp"]},
		}
	*/
	sv, err := slack.NewSecretsVerifier(httpHd, sig)
	fmt.Println(err)

	sv.Write([]byte(request.Body))
	if er := sv.Ensure(); er != nil {
		return fmt.Errorf("can't verify sig")
	}
	return nil
}
