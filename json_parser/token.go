// what token exist in JSON parser 

package main


type TokenType int


// Create a group of constants 
// iota is a Go's constant generator 
// iota automatically increments the constant
const (
	
	// special tokens
	EOF TokenType=iota
	ILLEGAL

	// Delimiters
	LEFT_BRACE
	RIGHT_BRACE


	COLON
	COMMA
	STRING
)


type Token struct {
	Type TokenType
	Literal string
	Line int 
	Column int
}

