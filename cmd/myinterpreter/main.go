package main

import (
    "fmt"
    "os"

    "github.com/codecrafters-io/interpreter-starter-go/token"
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

    tokens, err := token.Tokenize(file)
    if err != nil {
        return fmt.Errorf("error tokenizing: %w", err)
    }

    containsError := false
    for _, t := range tokens {
        if !containsError {
            containsError = t.Type == token.Error
        }

        if t.Type == token.Error {
            fmt.Fprint(os.Stderr, t.String())
        } else {
            fmt.Fprint(os.Stdout, t.String())
        }
    }

    if containsError {
        os.Exit(65)
    } else {
        os.Exit(0)
    }

    return nil
}
