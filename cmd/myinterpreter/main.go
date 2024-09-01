package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type TokenType string

const (
	LeftParen  TokenType = "LEFT_PAREN"
	RightParen TokenType = "RIGHT_PAREN"
	LeftBrace  TokenType = "LEFT_BRACE"
	RightBrace TokenType = "RIGHT_BRACE"
	EOF        TokenType = "EOF"
)

type Token struct {
	Type   TokenType
	Lexeme string
}

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

	tokens, err := tokenize(file)
	if err != nil {
		return fmt.Errorf("error tokenizing: %w", err)
	}

	printTokens(tokens)
	return nil
}

func tokenize(r io.Reader) ([]Token, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var tokens []Token
	for scanner.Scan() {
		char := scanner.Text()
		switch char {
		case "(":
			tokens = append(tokens, Token{LeftParen, char})
		case ")":
			tokens = append(tokens, Token{RightParen, char})
		case "{":
			tokens = append(tokens, Token{LeftBrace, char})
		case "}":
			tokens = append(tokens, Token{RightBrace, char})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning input: %w", err)
	}

	tokens = append(tokens, Token{EOF, ""})
	return tokens, nil
}

func printTokens(tokens []Token) {
	for _, token := range tokens {
		fmt.Printf("%s %s null\n", token.Type, token.Lexeme)
	}
}
