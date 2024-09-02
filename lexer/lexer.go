package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/codecrafters-io/interpreter-starter-go/token"
)

func Tokenize(r io.Reader) ([]token.Token, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var tokens []token.Token
	lineNumber := 1

	for scanner.Scan() {
		lineText := scanner.Text()
		lineText = stripComments(lineText)

		lineTokens, err := tokenizeLine(lineText, lineNumber)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, lineTokens...)

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning input: %w", err)
	}

	tokens = append(tokens, token.Token{Type: token.EOF, Line: lineNumber})
	return tokens, nil
}

func tokenizeLine(line string, lineNumber int) ([]token.Token, error) {
	var tokens []token.Token
	reader := strings.NewReader(line)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)

lineLoop:
	for scanner.Scan() {
		char := scanner.Text()

		t := token.Token{Line: lineNumber}

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
		case "/":
			t.Type, t.Lexeme = token.Slash, char
		case "=":
			if len(tokens) > 0 {
				lastToken := &tokens[len(tokens)-1]
				for _, composite := range token.Composites {
					if lastToken.Type == composite.Base {
						lastToken.Type = composite.Full
						lastToken.Lexeme += char
						continue lineLoop
					}
				}
			}
			t.Type, t.Lexeme = token.Equal, char
		case "!":
			t.Type, t.Lexeme = token.Bang, char
		case " ", "\t":
			continue
		default:
			t.Type, t.Lexeme = token.Error, char
		}

		tokens = append(tokens, t)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning line: %w", err)
	}

	return tokens, nil
}

func stripComments(line string) string {
	if commentIndex := strings.Index(line, token.BEGINNING_OF_COMMENT); commentIndex != -1 {
		return line[:commentIndex]
	}
	return line
}
