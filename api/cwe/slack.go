package main

/*
func verifySig(request events.APIGatewayProxyRequest) error {
	httpHeader := make(map[string][]string)
	for k, v := range request.Headers {
		httpHeader[k] = []string{v}
	}

	sv, err := slack.NewSecretsVerifier(httpHeader, sigsec)
	if err != nil {
		fmt.Println(err)
	}

	sv.Write([]byte(request.Body))
	if er := sv.Ensure(); er != nil {
		return fmt.Errorf("cant verify sig")
	}
	return nil
}
*/
