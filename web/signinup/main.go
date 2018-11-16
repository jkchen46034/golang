package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const hashCost = 8

var db *sql.DB

func main() {
	http.HandleFunc("/signin", logging(Signin))
	http.HandleFunc("/signup", logging(Signup))
	initDB()
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func initDB() {
	var err error
	db, err = sql.Open("postgres", "user=psqluser123 password=psqluser123 dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("databse ping success")
}
