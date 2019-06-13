#!/bin/env bash

. ../config

declare -A sec

# if this paramter don't set ,http:// https:// will fetch to remote
aws configure --profile default set cli_follow_urlparam false

# cat <<'EOF'> slack_secrets
# sec["slack_webhookurl"]=''
# sec["slack_signingsecret"]=''
# EOF

. ./slack_secrets


for k in ${!sec[@]}; do
  name="/${prefix}/${k}"
  echo $name
  echo ${sec[$k]}
  
  aws ssm put-parameter \
    --name $name \
    --value ${sec[$k]} \
    --type SecureString  > /dev/null
done

#for k in ${!sec[@]};do
#  aws ssm get-parameters --name "/${prefix}/${sec[$k]}" --with-decryption

#done
