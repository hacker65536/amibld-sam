package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)

var (
	cfg aws.Config
)

func init() {
	cfg, _ = external.LoadDefaultAWSConfig()
}
