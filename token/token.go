package token

import (
	"fmt"
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
	Error      Type = "Error"
	Equal      Type = "EQUAL"
	EqualEqual Type = "EQUAL_EQUAL"
)

type Token struct {
	Type   Type
	Lexeme string
	Line   int
}

func (t Token) String() string {
	if t.Type == Error {
		return fmt.Sprintf("[line %d] Error: Unexpected character: %s", t.Line, t.Lexeme)
	}
	return fmt.Sprintf("%s %s null", t.Type, t.Lexeme)
}
