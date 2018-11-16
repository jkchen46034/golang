package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	//"strings"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Println(r.Form)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//fmt.Println("key:", k)
	//fmt.Println("val:", strings.Join(v, ""))
	//}
	fmt.Fprintf(w, "Hello China!")
}

func login(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", logging(sayhelloName))
	http.HandleFunc("/login", logging(login))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
