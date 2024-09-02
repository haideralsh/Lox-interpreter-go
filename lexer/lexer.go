package lexer

import (
	"bufio"
	"fmt"
	"io"

	"github.com/codecrafters-io/interpreter-starter-go/token"
)

func Tokenize(r io.Reader) ([]token.Token, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var tokens []token.Token
	line := 1

	for scanner.Scan() {
		char := scanner.Text()

		t := token.Token{Line: line}

		switch char {
		case "(":
			t.Type, t.Lexeme = token.LeftParen, char
		case ")":
			t.Type, t.Lexeme = token.RightParen, char
		case "{":
			t.Type, t.Lexeme = token.LeftBrace, char
		case "}":
			t.Type, t.Lexeme = token.RightBrace, char
		case ",":
			t.Type, t.Lexeme = token.Comma, char
		case ".":
			t.Type, t.Lexeme = token.Dot, char
		case "-":
			t.Type, t.Lexeme = token.Minus, char
		case "+":
			t.Type, t.Lexeme = token.Plus, char
		case ";":
			t.Type, t.Lexeme = token.SemiColon, char
		case "*":
			t.Type, t.Lexeme = token.Star, char
		case "<":
			t.Type, t.Lexeme = token.Less, char
		case ">":
			t.Type, t.Lexeme = token.Greater, char
		case "=":
			if len(tokens) > 0 {
				switch tokens[len(tokens)-1].Type {
				case token.Equal:
					tokens[len(tokens)-1].Type = token.EqualEqual
					tokens[len(tokens)-1].Lexeme += char
					continue
				case token.Bang:
					tokens[len(tokens)-1].Type = token.BangEqual
					tokens[len(tokens)-1].Lexeme += char
					continue
				case token.Less:
					tokens[len(tokens)-1].Type = token.LessEqual
					tokens[len(tokens)-1].Lexeme += char
					continue
				case token.Greater:
					tokens[len(tokens)-1].Type = token.GreaterEqual
					tokens[len(tokens)-1].Lexeme += char
					continue
				}
			}
			t.Type, t.Lexeme = token.Equal, char
		case "!":
			t.Type, t.Lexeme = token.Bang, char
		case "\n":
			line++
			continue
		default:
			t.Type, t.Lexeme = token.Error, char
		}

		tokens = append(tokens, t)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning input: %w", err)
	}

	tokens = append(tokens, token.Token{Type: token.EOF, Line: line})
	return tokens, nil
}
