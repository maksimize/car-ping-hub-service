package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/go-redis/redis"

)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	HTTPMethodNotSupported = errors.New("no name was provided in the HTTP body")
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("Body size = %d. \n", len(request.Body))
	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("  %s: %s\n", key, value)
	}
	if request.HTTPMethod == "GET" {
		fmt.Printf("GET METHOD\n")
		return events.APIGatewayProxyResponse{Body: "GET", StatusCode: 200}, nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}
