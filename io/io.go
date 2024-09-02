package io

import (
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/token"
)

func Print(tokens []token.Token) {
	for _, t := range tokens {
		t.Print()
	}
}

func Exit(tokens []token.Token) {
	for _, t := range tokens {
		if t.Type == token.Error {
			os.Exit(65)
		}
	}
	os.Exit(0)
}
