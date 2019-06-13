#!/bin/env bash

. ./config

opt="${profile:+--profile ${profile}} ${region:+--region ${region}}"

aws ${opt} ssm get-parameters-by-path --path /${prefix}/slacksigsec --recursive | jq '.[][]|.Name' | xargs -I{} aws ${opt} ssm delete-parameter --name {}
aws ${opt} ssm get-parameters-by-path --path /${prefix}/slackoauthtoken --recursive | jq '.[][]|.Name' | xargs -I{} aws ${opt} ssm delete-parameter --name {}
aws ${opt} ssm get-parameters-by-path --path /${prefix}/slackchannel --recursive | jq '.[][]|.Name' | xargs -I{} aws ${opt} ssm delete-parameter --name {}
aws ${opt} ssm get-parameters-by-path --path /${prefix}/slackwebhook --recursive | jq '.[][]|.Name' | xargs -I{} aws ${opt} ssm delete-parameter --name {}

