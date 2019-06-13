#!/bin/env bash

set -e 
conf=config

#{{{
NOCOLOR='\033[0m'
RED='\033[0;31m'
GREEN='\033[0;32m'
ORANGE='\033[0;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
LIGHTGRAY='\033[0;37m'
DARKGRAY='\033[1;30m'
LIGHTRED='\033[1;31m'
LIGHTGREEN='\033[1;32m'
YELLOW='\033[1;33m'
LIGHTBLUE='\033[1;34m'
LIGHTPURPLE='\033[1;35m'
LIGHTCYAN='\033[1;36m'
WHITE='\033[1;37m'
#}}}

function s3mb(){
	echo -e "Bucket doesn't exist or not owned.\n${PURPLE}create bucket?${NOCOLOR}"
	read -p '(y/n)' key

	if [[ "$key" != "y" ]]; then
		echo "aborted"
		return	
	fi
	opt="${profile:+--profile ${profile}} ${region:+--region ${region}}"
	aws ${opt} s3 mb s3://${s3} 
}
if [[ ! -e $conf ]];then

	echo -e "input info for ${RED}sam${NOCOLOR}\n"
	read -p 'profile(aws)                      :' profile
	read -p 'region(aws)                       :' region
	read -p 'prefix(same as tf workspace)      :' prefix
	read -p 's3(<prefix>-sam)                  :' s3

	
        chk=${prefix:? no input}

	echo 
	echo
	echo "---confirm----"
	echo "profile   :${profile:-(default)}"
	defregion=$(aws configure get region)
	echo "region    :${region:-${defregion}}"
	echo "prefix    :$prefix"
	echo "s3        :${s3:=${prefix}-sam}"
	echo "--------------"


	read -p 'init (y/n)' key
	if [[ "$key" != "y" ]]; then
		echo "aborted"
		exit 0
	fi

fi

aws s3api head-bucket --bucket $s3 >/dev/null 2>&1 && \
	echo -e "${GREEN}Bucket exists${NOCOLOR}\n" || \
	s3mb

cat <<EOF > $conf
prefix=${prefix}
s3=$s3
profile=${profile}
region=${region}
EOF
echo -e "${GREEN}wrote to $conf${NOCOLOR}\n"




# vim:set foldmethod=marker: #
