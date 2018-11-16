package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Credentials struct {
	Username string `json:"username", db:"username"`
	Password string `json:"password", db:"password"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("json decoding failed")
		return
	}
	fmt.Println(creds)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	fmt.Println("password: ", string(hashedPassword))
	if _, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("database write error", err.Error())
		return
	}
	// default status of 200 is sent back
}

func Signin(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("json decode failed")
		return
	}
	result := db.QueryRow("select password from users where username=$1", creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("database query failed")
		return
	}
	storedCreds := &Credentials{}
	err = result.Scan(&storedCreds.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("login failed, password does not match")
		return
	}

	// The default 200 status is sent
	fmt.Println("matched and signed in")
}
