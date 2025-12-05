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

	p.nextToken() // read past the next token

	if p.currentToken.Type == RIGHT_BRACE {
		
		p.nextToken()

		if p.currentToken.Type != EOF {
			fmt.Printf("Error: Expected end of input but got '%s' at line no %d, at column no %d", 
				p.currentToken.Literal, p.currentToken.Line, p.currentToken.Column)
			return false
		}

		return true
	}
		
	fmt.Printf("Error: Object with content not implemented yet\n")
	return false
}

// Let's do the validation alone.
func (p *Parser) parseKeyValuePair() bool {
	
	if p.currentToken.Type != STRING {
		fmt.Printf("Error: Expected string key but got '%s' at line %d, column %d\n",
			p.currentToken.Literal, p.currentToken.Line, p.currentToken.Column)
		return false
	}
	
	p.nextToken()

	if p.currentToken.Type != COLON {
		fmt.Printf("Error: Expected ':' but got '%s' at line %d, column %d\n",
			p.currentToken.Literal, p.currentToken.Line, p.currentToken.Column)
		return false
	}

	p.nextToken()

	if p.currentToken.Type != STRING {
		fmt.Printf("Error: Expected string value but got '%s' at line %d, column %d\n",
			p.currentToken.Literal, p.currentToken.Line, p.currentToken.Column)
		return false
	}

	p.nextToken()
	return true
}


