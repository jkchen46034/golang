package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const hashCost = 8

var db *sql.DB

func main() {
	// "Signin" and "Signup" are handler that we will implement
	http.HandleFunc("/signin", logging(Signin))
	http.HandleFunc("/signup", logging(Signup))
	// initialize our database connection
	initDB()
	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func initDB() {
	var err error
	// Connect to the postgres db
	//you might have to change the connection string to add your database credentials
	//db, err = sql.Open("postgres", "user=psqluser123 password=psqluser123 dbname=mydb sslmode=disable")
	db, err = sql.Open("postgres", "user=psqluser123 password=psqluser123 dbname=mydb sslmode=disable")
	if err != nil {
		panic(err)
	}
}
