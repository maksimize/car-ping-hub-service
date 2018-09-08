package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"gopkg.in/redis.v3"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
})
var port = 8200

type Response struct {
	Status string
}
func handle(w http.ResponseWriter, r *http.Request) {
	client.Publish("radio", "test MAKS")

	//message := r.URL.Path
	//fmt.Println(message)
	//message = strings.TrimPrefix(message, "/")

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	response := Response{"success"}

	jsonResponse, err := json.Marshal(response)
	if err != nil{
		panic(err)
	}
	w.Write(jsonResponse)
}

func main() {
	fmt.Println("====App Starts====")
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
