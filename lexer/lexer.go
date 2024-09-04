package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
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
		foundNonNumber := false

		if isNumber(char) {
			literal := []string{char}
			for scanner.Scan() {
				char = scanner.Text()
				if isNumber(char) {
					literal = append(literal, char)
				} else {
					foundNonNumber = true
					break
				}
			}

			strLiteral := strings.Join(literal, "")

			if every(literal, isDot) {
				for _, l := range literal {
					tokens = append(tokens, token.Token{Type: token.Dot, Lexeme: l})
				}
			} else {
				tokens = append(tokens, token.Token{Type: token.Number, Lexeme: strLiteral, Literal: fmtNumberLiteral(strLiteral)})
			}

			if !foundNonNumber {
				break
			}
		}

		if is(char, "(") {
			tokens = append(tokens, token.Token{Type: token.LeftParen, Lexeme: char})
		} else if is(char, ")") {
			tokens = append(tokens, token.Token{Type: token.RightParen, Lexeme: char})
		} else if is(char, "{") {
			tokens = append(tokens, token.Token{Type: token.LeftBrace, Lexeme: char})
		} else if is(char, "}") {
			tokens = append(tokens, token.Token{Type: token.RightBrace, Lexeme: char})
		} else if is(char, ",") {
			tokens = append(tokens, token.Token{Type: token.Comma, Lexeme: char})
		} else if is(char, ".") {
			tokens = append(tokens, token.Token{Type: token.Dot, Lexeme: char})
		} else if is(char, "-") {
			tokens = append(tokens, token.Token{Type: token.Minus, Lexeme: char})
		} else if is(char, "+") {
			tokens = append(tokens, token.Token{Type: token.Plus, Lexeme: char})
		} else if is(char, ";") {
			tokens = append(tokens, token.Token{Type: token.SemiColon, Lexeme: char})
		} else if is(char, "*") {
			tokens = append(tokens, token.Token{Type: token.Star, Lexeme: char})
		} else if is(char, "<") {
			tokens = append(tokens, token.Token{Type: token.Less, Lexeme: char})
		} else if is(char, ">") {
			tokens = append(tokens, token.Token{Type: token.Greater, Lexeme: char})
		} else if is(char, "/") {
			for scanner.Scan() {
				char = scanner.Text()
				if char == "/" {
					break lineLoop
				}
			}

			tokens = append(tokens, token.Token{Type: token.Slash, Lexeme: char})
		} else if is(char, "=") {
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
		} else if is(char, "!") {
			tokens = append(tokens, token.Token{Type: token.Bang, Lexeme: char})
		} else if is(char, " ") || is(char, "\t") {
			continue
		} else if is(char, "\"") {
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
		} else {
			errors = append(errors, token.Error{Line: lineNumber, Message: fmt.Sprintf("Unexpected character: %s", char)})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error scanning line: %w", err)
	}

	return tokens, errors, nil
}

func is(a, b string) bool {
	return a == b
}

func isNumber(char string) bool {
	if isDigit(char) {
		return true
	}

	if isDot(char) {
		return true
	}

	return false
}

func isDigit(char string) bool {
	return char >= "0" && char <= "9"
}

func isDot(char string) bool {
	return char == "."
}

func fmtNumberLiteral(literal string) string {
	floatLiteral, _ := strconv.ParseFloat(literal, 64)

	if !strings.Contains(literal, ".") {
		return fmt.Sprintf("%.1f", floatLiteral)
	}

	return fmt.Sprintf("%.2f", float64(floatLiteral))
}

func every[T any](arr []T, condition func(T) bool) bool {
	for _, v := range arr {
		if !condition(v) {
			return false
		}
	}
	return true
}
