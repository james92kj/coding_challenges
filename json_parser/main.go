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

	// Get the filename from command line arguments 
	filename := os.Args[1]

	// Read the file 
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Parse the Json 
	parser := NewParser(string(content))

	if parser.Parse() {
		fmt.Println("\nValid JSON")
	} else {
		fmt.Println("\nInvalid JSON")
	}

}


