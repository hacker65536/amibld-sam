#!/bin/env bash

docker pull amazon/dynamodb-local
docker network create lambda-local


# if do docker run before mkdir data directory , its owner will be root and dynamdob cant access dbfile
mkdir ./data

docker-compose up
#docker run --rm -d -v $(pwd)/data:/data -p 8000:8000 --network lambda-local --name dynamodb amazon/dynamodb-local -jar DynamoDBLocal.jar -dbPath /data -sharedDb
