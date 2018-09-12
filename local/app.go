package main

import (
	"fmt"
	"os"
	"net/http"
	"strings"
	"encoding/json"
	"gopkg.in/redis.v3"
)
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
func handle(w http.ResponseWriter, r *http.Request) {

	VehicleId := r.URL.Path
	VehicleId = strings.TrimPrefix(VehicleId, endPoint)
	fmt.Println(VehicleId)

	client.Publish(redisChannel, VehicleId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := Response{"success"}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Write(jsonResponse)
}

func main() {
	fmt.Println("====Car Ping Hub Service App Starts====")
	http.HandleFunc(endPoint, handle)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
