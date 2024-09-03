package output

import (
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/token"
)

func Print(tokens []token.Token, errors []token.Error) {
	for _, e := range errors {
		e.Print()
	}
	for _, t := range tokens {
		t.Print()
	}
}

func Exit(errors []token.Error) {
	if len(errors) > 0 {
		os.Exit(65)
	}

	os.Exit(0)

}
