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
}
