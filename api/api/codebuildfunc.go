package main

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
)

var (
	codebuildsvc *codebuild.Client
)

func init() {
	codebuildsvc = codebuild.New(cfg)
}

func listBuildsRequest() *codebuild.ListBuildsResponse {

	input := &codebuild.ListBuildsInput{}
	req := codebuildsvc.ListBuildsRequest(input)
	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

func batchGetProjectsRequest() *codebuild.BatchGetProjectsResponse {

	input := &codebuild.BatchGetProjectsInput{
		Names: []string{
			"amibld-codebuild-centos6",
		},
	}
	req := codebuildsvc.BatchGetProjectsRequest(input)
	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

func startBuildRequest(be *Buildcfg) *codebuild.StartBuildResponse {

	input := &codebuild.StartBuildInput{}

	input.ProjectName = aws.String(prefix + "-codebuild-" + be.Platform)

	if be.UserName != "" {
		input.EnvironmentVariablesOverride = append(input.EnvironmentVariablesOverride, codebuild.EnvironmentVariable{
			Name:  aws.String("USER_NAME"),
			Value: aws.String(be.UserName),
		})
	}

	if be.AMI != "" {
		input.EnvironmentVariablesOverride = append(input.EnvironmentVariablesOverride, codebuild.EnvironmentVariable{
			Name:  aws.String("SOURCE_AMI"),
			Value: aws.String(be.AMI),
		})
	}

	if be.DynamoTable != "" {
		input.EnvironmentVariablesOverride = append(input.EnvironmentVariablesOverride, codebuild.EnvironmentVariable{
			Name:  aws.String("DYNAMOTABLE"),
			Value: aws.String(be.DynamoTable),
		})
	}

	if len(be.Account) != 0 {
		input.EnvironmentVariablesOverride = append(input.EnvironmentVariablesOverride, codebuild.EnvironmentVariable{
			Name:  aws.String("AMI_USERS"),
			Value: aws.String(strings.Join(be.Account, ",")),
		})
	}

	if len(be.Region) != 0 {
		input.EnvironmentVariablesOverride = append(input.EnvironmentVariablesOverride, codebuild.EnvironmentVariable{
			Name:  aws.String("AMI_REGIONS"),
			Value: aws.String(strings.Join(be.Region, ",")),
		})
	}

	if len(be.Role) != 0 {
		input.EnvironmentVariablesOverride = append(input.EnvironmentVariablesOverride, codebuild.EnvironmentVariable{
			Name:  aws.String("ROLES"),
			Value: aws.String(strings.Join(be.Role, ",")),
		})
	}

	if be.Commit != "" {
		input.SourceVersion = aws.String(be.Commit)

	}

	fmt.Println("input:", input)
	//resp := &codebuild.StartBuildResponse{}
	req := codebuildsvc.StartBuildRequest(input)
	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

func searchValue(name string, env []codebuild.EnvironmentVariable) string {

	sort.Slice(env, func(i, j int) bool {
		return *env[i].Name <= *env[j].Name
	})

	n := name
	idx := sort.Search(len(env), func(i int) bool {
		return string(*env[i].Name) >= n
	})

	if *env[idx].Name == n {
		return *env[idx].Value
	}
	return ""
}
