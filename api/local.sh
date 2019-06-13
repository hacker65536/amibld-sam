#!/bin/env bash

make build && \
sam local start-api -p 8080 --docker-network lambda-local --env-vars env.json   --host $(hostname -I | cut -d ' ' -f1) --debug
#sam local start-api -p 8080 --docker-network lambda-local  --host $(hostname -I | cut -d ' ' -f1) --debug

