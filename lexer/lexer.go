package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/codecrafters-io/interpreter-starter-go/token"
)

func Tokenize(r io.Reader) ([]token.Token, []token.Error, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var tokens []token.Token
	var errors []token.Error

	lineNumber := 1

	for scanner.Scan() {
		lineText := scanner.Text()

		lineTokens, lineErrors, err := tokenizeLine(lineText, lineNumber)
		if err != nil {
			return nil, nil, err
		}
		tokens = append(tokens, lineTokens...)
		errors = append(errors, lineErrors...)

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error scanning input: %w", err)
	}

	tokens = append(tokens, token.Token{Type: token.EOF, Line: lineNumber})

	return tokens, errors, nil
}

func tokenizeLine(line string, lineNumber int) ([]token.Token, []token.Error, error) {
	var tokens []token.Token
	var errors []token.Error

	reader := strings.NewReader(line)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)

lineLoop:
	for scanner.Scan() {
		char := scanner.Text()

		switch char {
		case "(":
			tokens = append(tokens, token.Token{Type: token.LeftParen, Lexeme: char})
		case ")":
			tokens = append(tokens, token.Token{Type: token.RightParen, Lexeme: char})
		case "{":
			tokens = append(tokens, token.Token{Type: token.LeftBrace, Lexeme: char})
		case "}":
			tokens = append(tokens, token.Token{Type: token.RightBrace, Lexeme: char})
		case ",":
			tokens = append(tokens, token.Token{Type: token.Comma, Lexeme: char})
		case ".":
			tokens = append(tokens, token.Token{Type: token.Dot, Lexeme: char})
		case "-":
			tokens = append(tokens, token.Token{Type: token.Minus, Lexeme: char})
		case "+":
			tokens = append(tokens, token.Token{Type: token.Plus, Lexeme: char})
		case ";":
			tokens = append(tokens, token.Token{Type: token.SemiColon, Lexeme: char})
		case "*":
			tokens = append(tokens, token.Token{Type: token.Star, Lexeme: char})
		case "<":
			tokens = append(tokens, token.Token{Type: token.Less, Lexeme: char})
		case ">":
			tokens = append(tokens, token.Token{Type: token.Greater, Lexeme: char})
		case "/":
			for scanner.Scan() {
				char = scanner.Text()
				if char == "/" {
					break lineLoop
				}
			}

			tokens = append(tokens, token.Token{Type: token.Slash, Lexeme: char})
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
			tokens = append(tokens, token.Token{Type: token.Equal, Lexeme: char})
		case "!":
			tokens = append(tokens, token.Token{Type: token.Bang, Lexeme: char})
		case " ", "\t":
			continue
		case "\"":
			literal := ""
			for scanner.Scan() {
				char = scanner.Text()
				if char == "\"" {
					break
				}
				literal += char
			}
			if char != "\"" {
				errors = append(errors, token.Error{Line: lineNumber, Message: "Unterminated string."})
			} else {
				lexeme := fmt.Sprintf("\"%s\"", literal)
				tokens = append(tokens, token.Token{Type: token.String, Lexeme: lexeme, Literal: literal})
			}
		default:
			errors = append(errors, token.Error{Line: lineNumber, Message: fmt.Sprintf("Unexpected character: %s", char)})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error scanning line: %w", err)
	}

	return tokens, errors, nil
}
