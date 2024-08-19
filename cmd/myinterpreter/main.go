package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

    if len(os.Args) < 3 {
        fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
        os.Exit(1)
    }

    command := os.Args[1]

    if command != "tokenize" {
        fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
        os.Exit(1)
    }

    filename := os.Args[2]
    fileContents, err := os.ReadFile(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
        os.Exit(1)
    }

    if len(fileContents) > 0 {
        var slice []string

        for _, v := range fileContents {
            switch string(v) {
            case "(":
                slice = append(slice, "LEFT_PAREN ( null")
            case ")":
                slice = append(slice, "RIGHT_PAREN ) null")
            case "{":
                slice = append(slice, "LEFT_BRACE { null")
            case "}":
                slice = append(slice, "RIGHT_BRACE } null")
            }
        }

        slice = append(slice, "EOF  null")

        fmt.Println(strings.Join(slice, "\n"))
    } else {
        fmt.Println("EOF  null")
    }
}
