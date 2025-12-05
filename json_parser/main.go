package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: json-parser <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	parser := NewParser(string(content))

	if parser.Parse() {
		fmt.Println("Valid JSON")
		os.Exit(0)
	} else {
		fmt.Println("Invalid JSON")
		os.Exit(1)
	}
}
