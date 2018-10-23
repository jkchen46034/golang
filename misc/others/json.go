package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Human struct {
		ID       int
		NickName string
		Alias    []string
	}

	human := Human{
		ID:       1,
		NickName: "J",
		Alias:    []string{"Jerry Foulton", "Jaja Mango", "Joking Zhang"},
	}
	b, err := json.Marshal(human)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)
	fmt.Println()

	jsonblob := []byte(`{"ID":2, "Nickname": "Katy", "Alias":["Mr. JacK", "Miss Ki", "Joking K"]}`)
	var jsonobj Human
	err = json.Unmarshal(jsonblob, &jsonobj)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(jsonobj)

	type T struct {
		F1 int `json:"f_1"`
		F2 int `json:"f_2,omitempty"`
		F3 int `json:"f_3,omitempty"`
		F4 int `json:"-"`
	}
	t := T{1, 0, 2, 3}
	b, err = json.Marshal(t)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s\n", b)
	var t1 T
	err = json.Unmarshal(b, &t1)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(t1)

}

/*
$ go run json.go
{"ID":1,"NickName":"J","Alias":["Jerry Foulton","Jaja Mango","Joking Zhang"]}
{2 Katy [Mr. JacK Miss Ki Joking K]}
{"f_1":1,"f_3":2}
{1 0 2 0}
*/
