#!/bin/env bash

set -e

## set to .aws/config 
## cli_follow_urlparam = false

. ./config





opt="${profile:+--profile ${profile}} ${region:+--region ${region}}"

function inputsec(){
	read -p    'slack app  (appid)       :' ap
	read -s -p 'slack signing secret     :' ss
	echo
	read -s -p 'slack oauthtoken         :' token 
	echo
	read -p    'slack channelID          :' channel
	read -p    'slack webhookurl         :' webhookurl
	read -p    'tag                      :' tag


	echo
	set -f
	echo 
	echo
	echo "------------confirm--------------"
	echo "app        :$ap"
	xx=${ss:6:-6}
	xcnt=${#xx}
	xxx=$(eval printf '*%.s' {1..$xcnt}}|echo $(cat))
	echo "secret     :${ss: 0:5}${xxx}${ss: -5}"
	xx=${token:11:-6}
	xcnt=${#xx}
	xxx=$(eval printf '*%.s' {1..$xcnt}}|echo $(cat))
	echo "token      :${token: 0:10}${xxx}${token: -5}"
	echo "channelID  :$channel"
	echo "webhookurl :$webhookurl"
	echo "tag        :$tag"
	echo "---------------------------------"

	aws ${opt} ssm put-parameter \
		--name "/${prefix}/slacksigsec/$ap" \
		--value "${ss}" \
		--type SecureString  \
		--overwrite && \
		aws ${opt} ssm add-tags-to-resource \
		--resource-type Parameter \
		--resource-id "/${prefix}/slacksigsec/$ap" \
		--tags Key=ws,Value=$tag
	aws ${opt} ssm put-parameter \
		--name "/${prefix}/slackoauthtoken/$ap" \
		--value "${token}" \
		--type SecureString  \
		--overwrite && \
		aws ${opt} ssm add-tags-to-resource \
		--resource-type Parameter \
		--resource-id "/${prefix}/slackoauthtoken/$ap" \
		--tags Key=ws,Value=$tag
	aws ${opt} ssm put-parameter \
		--name "/${prefix}/slackwebhook/$ap" \
		--value "${webhookurl}" \
		--type SecureString  \
		--overwrite && \
		aws ${opt} ssm add-tags-to-resource \
		--resource-type Parameter \
		--resource-id "/${prefix}/slackwebhook/$ap" \
		--tags Key=ws,Value=$tag
	aws ${opt} ssm put-parameter \
		--name "/${prefix}/slackchannel/$ap" \
		--value "${channel}" \
		--type SecureString  \
		--overwrite && \
		aws ${opt} ssm add-tags-to-resource \
		--resource-type Parameter \
		--resource-id "/${prefix}/slackchannel/$ap" \
		--tags Key=ws,Value=$tag
}


inputsec

while read -s -p 'add another (y/n)' key
do
	if [[ $key != "y" ]];then
		break
	fi
	echo
	inputsec
done
echo
echo "done"

<<COMMENT
aws ${opt} ssm put-parameter \
  --name "/${prefix}/slackoauthtoken/$1" \
  --value "$3" \
  --type SecureString 


aws ${opt} ssm get-parameter --name "/${prefix}/slackoauthtoken/$1"
COMMENT
