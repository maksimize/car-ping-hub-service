package main

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/redis.v3"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)
var redisChannel = "car_status"
var redisHost = "car-status-channel.nybflb.0001.apse1.cache.amazonaws.com:6379"
var client = redis.NewClient(&redis.Options{
	Addr: redisHost,
})
var (
	// ErrNameNotProvided is thrown when a name is not provided
	HTTPMethodNotSupported = errors.New("no name was provided in the HTTP body")
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "GET" {
		var VehicleId = request.QueryStringParameters["vid"]
		fmt.Printf("=> vid is " + VehicleId)
		fmt.Printf("=> redis host " + redisHost)
		client.Publish(redisChannel, VehicleId)
		fmt.Printf("=> vid published " + VehicleId + " to " + redisHost)
		return events.APIGatewayProxyResponse{Body: VehicleId, StatusCode: 200}, nil
	} else {
		fmt.Printf("NEITHER\n")
		return events.APIGatewayProxyResponse{}, HTTPMethodNotSupported
	}
}

func main() {
	lambda.Start(HandleRequest)
}
