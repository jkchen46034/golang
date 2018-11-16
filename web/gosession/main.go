package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

var cache redis.Conn

func main() {
	initCache()
	http.HandleFunc("/signin", logging(Signin))
	http.HandleFunc("/welcome", logging(Welcome))
	http.HandleFunc("/refresh", logging(Refresh))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func initCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		log.Fatal(err)
	}
	cache = conn
	log.Println("redis connected")
}
