package main

import "fmt"

func main() {
	testCases := []struct {
		input string 
		expected bool
	}{
		{`{}`, true},
		{`{"name":"James"}`, true},
		{`{`,false},
		{`{"name"}`, false},
		{`{"name":}`, false},
	}

	for _, tc := range testCases {
		
		parser := NewParser(tc.input)
		is_valid := parser.Parse()

		if is_valid == tc.expected {
			fmt.Printf("Pass")
		} else {
			fmt.Printf("Not Pass")
		}
	}
}
