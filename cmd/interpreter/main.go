package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/lexer"
	"github.com/codecrafters-io/interpreter-starter-go/printer"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) != 3 || os.Args[1] != "tokenize" {
		return fmt.Errorf("usage: %s tokenize <filename>", os.Args[0])
	}

	file, err := os.Open(os.Args[2])
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	tokens, err := lexer.Tokenize(file)
	if err != nil {
		return fmt.Errorf("error tokenizing: %w", err)
	}

	printer.PrintAndTerminate(tokens)

	return nil
}
