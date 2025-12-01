package main

import "fmt"

type Parser struct{
	
	lexer *Lexer
	currentToken Token
	peekToken Token
}


func NewParser(input string) *Parser {
	
	p := &Parser{
		lexer: NewLexer(input),
	}

	p.nextToken()
	p.nextToken()
	
	return p
}


func (p *Parser) nextToken() {
	
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) Parse() bool {

	if p.currentToken.Type != LEFT_BRACE {
		
		fmt.Printf("Error: Expected '{' but got '%s' at line %d, column %d", p.currentToken.Literal, p.currentToken.Line, p.currentToken.Column)
		return false 
	}

	for p.currentToken.Type == ILLEGAL{
		fmt.Printf("\n Illegal Token %s",p.currentToken.Literal)
		p.nextToken()
	}

	if p.currentToken.Type != RIGHT_BRACE {
		fmt.Printf("Error: Expected '}' but got '%s' at line no %d, at column no %d", p.currentToken.Literal, p.currentToken.Line, p.currentToken.Column)
		return false
	}
	
	p.nextToken()
	if p.currentToken.Type != EOF {
		fmt.Printf("Error: Expected end of input but got '%s' at line no %d, at column no %d", p.currentToken.Literal, p.currentToken.Line, p.currentToken.Column)
		return false
	}

	return true
}
