package main

import (
    "flag"
    "fmt"
    "os"
    "bytes" 
    "bufio"
    "unicode/utf8"
    "io"
)

func main() {
    
    // Define the flags 
    countBytes := flag.Bool("c", false, "count bytes")
    countLines := flag.Bool("l", false, "count Lines")
    countWords := flag.Bool("w", false, "count words")
    countChars := flag.Bool("m", false, "count characters")


    // Parse the arguments 
    flag.Parse()
    
    noFlags := !*countBytes && !*countLines && !*countWords && !*countChars
    if noFlags {
        *countBytes = true
        *countLines = true 
        *countWords = true 
    }
    
    var data []byte
    var err error
    var filename string

    // Get the filename 
    args := flag.Args()
    
    if len(args) < 1 {
        data, err = io.ReadAll(os.Stdin)
        filename = ""
    } else {
        filename = args[0]
        data, err = os.ReadFile(filename)
    }
    
    
    if err != nil {
        fmt.Fprintf(os.Stderr, "ccwc: %v\n", err)
        os.Exit(1)
    }
    
    if *countBytes {
        byteCount := len(data)
        fmt.Printf("%7d",byteCount)
    }

    if *countLines {
        lineCount := bytes.Count(data, []byte{'\n'})
        fmt.Printf("%7d",lineCount)
    }
    
    if *countWords {
        wordCount := 0
        scanner := bufio.NewScanner(bytes.NewReader(data))
        scanner.Split(bufio.ScanWords)
        
        for scanner.Scan() {
            wordCount++
        }
        fmt.Printf("%7d",wordCount)

    }
    
    if *countChars {
        characterCount := utf8.RuneCount(data)
        fmt.Printf("%7d", characterCount)
    }

    if filename != "" {
       fmt.Printf(" %s\n", filename)
    } else {
       fmt.Println("")
    }
}


