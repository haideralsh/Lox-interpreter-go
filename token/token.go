package token

import (
	"bufio"
	"fmt"
	"io"
)

type Type string

const (
	LeftParen  Type = "LEFT_PAREN"
	RightParen Type = "RIGHT_PAREN"
	LeftBrace  Type = "LEFT_BRACE"
	RightBrace Type = "RIGHT_BRACE"
	EOF        Type = "EOF"
	Comma      Type = "COMMA"
	Dot        Type = "DOT"
	Minus      Type = "MINUS"
	Plus       Type = "PLUS"
	SemiColon  Type = "SEMICOLON"
	Star       Type = "STAR"
)

type Token struct {
	Type   Type
	Lexeme string
}

func Tokenize(r io.Reader) ([]Token, error) {
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
		case ",":
			tokens = append(tokens, Token{Comma, char})
		case ".":
			tokens = append(tokens, Token{Dot, char})
		case "-":
			tokens = append(tokens, Token{Minus, char})
		case "+":
			tokens = append(tokens, Token{Plus, char})
		case ";":
			tokens = append(tokens, Token{SemiColon, char})
		case "*":
			tokens = append(tokens, Token{Star, char})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning input: %w", err)
	}

	tokens = append(tokens, Token{EOF, ""})
	return tokens, nil
}

func Print(tokens []Token) {
	for _, token := range tokens {
		fmt.Printf("%s %s null\n", token.Type, token.Lexeme)
	}
}
