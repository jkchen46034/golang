package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(num uint64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buf.Bytes()
}
