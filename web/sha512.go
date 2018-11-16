package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	secret := "password"
	s512 := sha512.New()
	s512.Write([]byte(secret))
	hash := hex.EncodeToString(s512.Sum(nil))
	fmt.Println(secret, hash)
}
