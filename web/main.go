package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        //fmt.Println(r.URL)
        HandleIndex(w, r)
    })

    fmt.Println("Starting Server... port number, 5678")
    log.Fatal(http.ListenAndServe(":5678", nil))
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    w.Write([]byte("Hello, World!"))
}
