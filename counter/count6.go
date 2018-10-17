package main

import (
	"fmt"
	"strings"
)

// create a new func type that returns the next token
// in a tokenized string.
type tokenizer func() (token string, ok bool)

// split splits `s` by `sep` separator and
// returns an iterator which walks over
// the tokens of `s` using a closure.
func split(s, sep string) tokenizer {

	// closure will use this context of
	// tokens, last and total variables.
	tokens, last := strings.Split(s, sep), 0

	// returning an anonymous func with
	// tokenizer func type.
	return func() (string, bool) {
		if len(tokens) == last {
			return "", false
		}

		// advance the iterator
		last = last + 1

		return tokens[last-1], true
	}
}

func main() {
	const sentence = "The quick brown fox jumps over the lazy dog"

	iter := split(sentence, " ")

	for token, ok := iter(); ok; token, ok = iter() {
		fmt.Println(token)
	}
}
