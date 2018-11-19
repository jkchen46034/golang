package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Message struct {
	EventName     string
	EventLocation string
	EventId       int64
	Speed         []float64
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello welcome to my site localhost 101")
	return
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
	http.HandleFunc("/", logging(rootHandler))
	http.HandleFunc("/post", logging(postJson))
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// template
	type Todo struct {
		Title string
		Done  bool
	}

	type TodoPageData struct {
		PageTitle string
		Todos     []Todo
	}

	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/todo", logging(func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	}))

	// another one, for form

	type ContactDetails struct {
		Email   string
		Subject string
		Message string
	}

	tmplform := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/contact", logging(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmplform.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		_ = details

		tmplform.Execute(w, struct{ Success bool }{true})
	}))

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
