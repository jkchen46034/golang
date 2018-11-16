package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("sign in, inputs not supplied")
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Unauthorized user rejected, ", creds.Username)
		return
	}

	uid, _ := uuid.NewV4()
	sessionToken := uid.String()
	log.Println(sessionToken, "created")
	_, err = cache.Do("SETEX", sessionToken, "120", creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("unable to write session to redis")
		return
	}
	log.Println("token saved to redis")

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		log.Println("not autorized use, rejected", err.Error())
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	response, err := cache.Do("GET", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("unable to get from redis", err.Error())
		return
	}
	if response == nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("token not found in cache, use is rejected.")
		return
	}
	log.Println("token is good, use of the url is allowed.")
	w.Write([]byte(fmt.Sprintf("Welcome %s!", response)))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	response, err := cache.Do("GET", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if response == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	newid, _ := uuid.NewV4()
	newSessionToken := newid.String()
	_, err = cache.Do("SETEX", newSessionToken, "120", fmt.Sprintf("%s", response))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = cache.Do("DEL", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
}
