#!/bin/env bash

. ./config

opt="${profile:+--profile ${profile}} ${region:+--region ${region}}"

aws ${opt} ssm get-parameters-by-path --path /${prefix}/slacksigsec --recursive
aws ${opt} ssm get-parameters-by-path --path /${prefix}/slackoauthtoken --recursive
aws ${opt} ssm get-parameters-by-path --path /${prefix}/slackchannel --recursive
aws ${opt} ssm get-parameters-by-path --path /${prefix}/slackwebhook --recursive
