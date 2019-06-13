#!/bin/bash

. ../config

make build && \
sam package \
  --template-file template.yaml \
  --output-template-file packaged.yaml \
  --s3-bucket $s3  && \
sam deploy \
  --template-file packaged.yaml \
  --stack-name ${prefix}-$(basename $(pwd)) \
  --parameter-overrides PROJECTPREFIX=${prefix} \
  --capabilities CAPABILITY_IAM

