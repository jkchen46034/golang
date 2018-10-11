package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	EventName     string
	EventLocation string
	EventId       int64
	Speed         []float64
}

func postJson(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	var m Message
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	res, _ := json.Marshal(m)
	fmt.Println(string(res))

	m1 := Message{"True Golf", "Top Golf Lehi", 1234, []float64{102.0, 88.9, 44.2, 38.6, 25.7}}
	json.NewEncoder(w).Encode(m1)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/post", postJson)
	http.HandleFunc("/", helloWorld)
	fmt.Println("Listening on localhost:9999")
	log.Fatal(http.ListenAndServe(":9999", nil))
}

/*
$ go run server.go
Listening on localhost:9999
$ go run client.go
.... server received
{"EventName":"True Golf","EventLocation":"West Jordan 7000","EventId":1001,"Speed":[21,18.9,45]}
----- client received
{"EventName":"True Golf","EventLocation":"Top Golf Lehi","EventId":1234,"Speed":[102,88.9,44.2,38.6,25.7]}
$ curl http://localhost:9999
Hello, world!
*/
