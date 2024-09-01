package printer

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/token"
)

func PrintAndTerminate(tokens []token.Token) {
	containsError := false
	for _, t := range tokens {
		if !containsError {
			containsError = t.Type == token.Error
		}

		if t.Type == token.Error {
			fmt.Fprintln(os.Stderr, t.String())
		} else {
			fmt.Fprintln(os.Stdout, t.String())
		}
	}

	if containsError {
		os.Exit(65)
	} else {
		os.Exit(0)
	}
}
