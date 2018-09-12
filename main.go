package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"gopkg.in/redis.v3"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var HTTPMethodNotSupported = errors.New("no name was provided in the HTTP body")
var port = getEnv(":" + "PORT", ":8200")
var redisChannel = getEnv("REDISCHANNEL", "car_status")
var endPoint = "/"
var client = redis.NewClient(&redis.Options{
	Addr: getEnv("REDISHOST", "redis") + ":" + getEnv("REDISPORT", "6379"),
})

type Response struct {
	Status string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	VehicleId := request.Path
	VehicleId = strings.TrimPrefix(VehicleId, endPoint)

	if request.HTTPMethod == "GET" {
		client.Publish(redisChannel, VehicleId)
		return events.APIGatewayProxyResponse{Body: "success", StatusCode: 200}, nil
	}
	return events.APIGatewayProxyResponse{Body: "wrong url", StatusCode: 200}, nil

}

func main() {
	fmt.Println("====Car Ping Hub Service App Starts2====")
	lambda.Start(HandleRequest)
}
