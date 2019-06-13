package main

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var (
	ssmsvc *ssm.Client
)

func init() {
	ssmsvc = ssm.New(cfg)
}
