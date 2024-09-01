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
    Error      Type = "Error"
)

type Token struct {
    Type   Type
    Lexeme string
    line   int
}

func Tokenize(r io.Reader) ([]Token, error) {
    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanRunes)

    var tokens []Token
    line := 1

    for scanner.Scan() {
        char := scanner.Text()

        switch char {
        case "(":
            tokens = append(tokens, Token{LeftParen, char, line})
        case ")":
            tokens = append(tokens, Token{RightParen, char, line})
        case "{":
            tokens = append(tokens, Token{LeftBrace, char, line})
        case "}":
            tokens = append(tokens, Token{RightBrace, char, line})
        case ",":
            tokens = append(tokens, Token{Comma, char, line})
        case ".":
            tokens = append(tokens, Token{Dot, char, line})
        case "-":
            tokens = append(tokens, Token{Minus, char, line})
        case "+":
            tokens = append(tokens, Token{Plus, char, line})
        case ";":
            tokens = append(tokens, Token{SemiColon, char, line})
        case "*":
            tokens = append(tokens, Token{Star, char, line})
        case "/n":
            line++
        default:
            tokens = append(tokens, Token{Error, char, line})
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error scanning input: %w", err)
    }

    tokens = append(tokens, Token{EOF, "", line})
    return tokens, nil
}

func (t Token) String() string {
    switch t.Type {
    case Error:
        return fmt.Sprintf("[line %v] Error: Unexpected character: %v\n", t.line, t.Lexeme)
    default:
        return fmt.Sprintf("%s %s null\n", t.Type, t.Lexeme)
    }
}
