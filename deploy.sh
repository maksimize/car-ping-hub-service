#!/usr/bin/env bash

echo "Compiling handlers..."

GOOS=linux go build -o main
zip deployment.zip main

echo "Built the go file."
echo "Starting deploy via CloudFormation"

aws cloudformation package --template-file template.yaml --output-template-file serverless-output.yaml \
  --s3-bucket car-ping-hub-service
aws cloudformation deploy --template-file serverless-output.yaml --stack-name HelloGoWebApp --capabilities CAPABILITY_IAM

echo "Done."
