package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Message struct {
	EventName     string
	EventLocation string
	EventId       int64
	Speed         []float64
}

func main() {
	m := Message{"True Golf", "West Jordan 7000", 1001, []float64{21.0, 18.9, 45.0}}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	res, err := http.Post("http://localhost:9999/post", "application/json; charset=utf-8", b)
	//res, err := http.Get("http://localhost:9999/post")
	if err != nil {
		log.Panic(err)
	}
	io.Copy(os.Stdout, res.Body)
}
