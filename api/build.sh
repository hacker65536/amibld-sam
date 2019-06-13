#!/bin/bash

. ../config
opt="${profile:+--profile ${profile}} ${region:+--region ${region}}"

make build && \
sam package ${opt} \
  --template-file template.yaml \
  --output-template-file packaged.yaml \
  --s3-bucket $s3  && \
sam deploy ${opt} \
  --template-file packaged.yaml \
  --stack-name ${prefix}-$(basename $(pwd)) \
  --parameter-overrides "PROJECTPREFIX=${prefix}" \
  --capabilities CAPABILITY_IAM

