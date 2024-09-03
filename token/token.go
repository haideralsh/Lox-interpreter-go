package token

import (
	"fmt"
	"os"
)

type Type string

const BEGINNING_OF_COMMENT = "//"

const (
	LeftParen    Type = "LEFT_PAREN"
	RightParen   Type = "RIGHT_PAREN"
	LeftBrace    Type = "LEFT_BRACE"
	RightBrace   Type = "RIGHT_BRACE"
	EOF          Type = "EOF"
	Comma        Type = "COMMA"
	Dot          Type = "DOT"
	Minus        Type = "MINUS"
	Plus         Type = "PLUS"
	SemiColon    Type = "SEMICOLON"
	Star         Type = "STAR"
	Equal        Type = "EQUAL"
	EqualEqual   Type = "EQUAL_EQUAL"
	Bang         Type = "BANG"
	BangEqual    Type = "BANG_EQUAL"
	Less         Type = "LESS"
	LessEqual    Type = "LESS_EQUAL"
	Greater      Type = "GREATER"
	GreaterEqual Type = "GREATER_EQUAL"
	Slash        Type = "SLASH"
	String       Type = "STRING"
)

type composite struct {
	Base Type
	Full Type
}

var Composites = []composite{
	{Equal, EqualEqual},
	{Bang, BangEqual},
	{Less, LessEqual},
	{Greater, GreaterEqual},
}

type TokenInterface interface {
	String() string
	Print()
}

type Token struct {
	Type    Type
	Lexeme  string
	Literal interface{}
	Line    int
}

type Error struct {
	Message string
	Line    int
	Lexeme  string
}

func (t Token) String() string {
	if t.Literal != nil {
		return fmt.Sprintf("%s %s %v", t.Type, t.Lexeme, t.Literal)
	} else {
		return fmt.Sprintf("%s %s null", t.Type, t.Lexeme)
	}
}

func (t Token) Print() {
	fmt.Fprintln(os.Stdout, t.String())
}

func (e Error) String() string {
	return fmt.Sprintf("[line %d] Error: %s", e.Line, e.Message)
}

func (e Error) Print() {
	fmt.Fprintln(os.Stderr, e.String())
}
