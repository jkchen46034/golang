package main

import (
	"fmt"
	"net/http"
)

func main() {
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/test", handleAPITest)

	usersMux := http.NewServeMux()
	usersMux.HandleFunc("/test", handleUsersTest)

	topMux := http.NewServeMux()
	topMux.Handle("/api/", http.StripPrefix("/api", apiMux))
	topMux.Handle("/users/", http.StripPrefix("/users", usersMux))

	http.ListenAndServe(":1234", topMux)
}

func handleAPITest(w http.ResponseWriter, req *http.Request) {
	// Called for /api/test
	fmt.Fprintf(w, "Hello from /api/test\n")
}

func handleUsersTest(w http.ResponseWriter, req *http.Request) {
	// Called for /users/test
	fmt.Fprintf(w, "Hello from /users/test\n")
}
