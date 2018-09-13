# Hello Go WebApp
This is an AWS Serverless web app built with [SAM](https://github.com/awslabs/serverless-application-model). The API Gateway currently supports GET that will trigger the same [Lambda function](main.go). The Lambda function is written in [Go](https://golang.org/).

# Usage
All of the below commands are bundled into the [deploy.sh](deploy.sh) script. Simply run `./deploy.sh` from the hello-go-webapp directory.

### TO BUILD
```
go get -d -v
GOOS=linux go build -o main
zip deployment.zip main
```

### TO DEPLOY
```
aws cloudformation package --template-file template.yaml --output-template-file serverless-output.yaml \
  --s3-bucket car-ping-hub-service
aws cloudformation deploy --template-file serverless-output.yaml --stack-name CarPingHub --capabilities CAPABILITY_IAM

```
