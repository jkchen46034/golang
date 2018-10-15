package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func savePage(page *Page) error {
	err := ioutil.WriteFile(page.Title, page.Body, 0666)
	return err
}

func loadPage(filename string) (*Page, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &Page{filename, body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := string(r.URL.Path[len("/view/"):]) + ".txt"
	page, err := loadPage(title)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(page.Body))
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
