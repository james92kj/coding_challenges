package main

type Lexer struct {
	
	input 			string
	position 		int 	// current position in input (points to current character)
	readPosition	int 	// current reading position in input (after current char)
	ch				byte    // current char under examination
	line 			int 	// current line number 
	column			int 	// current column number 
}


func NewLexer(input string) *Lexer {
	
	l := &Lexer{
		input: input,
		line: 1,
		column: 0,
	}
	
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
	l.column++
	
	if l.ch == '\n' {
		l.line++
		l.column = 0
	}
}

func (l *Lexer) skipWhitespace() {
	
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() Token {
	
	var tok Token

	l.skipWhitespace()

	tok.Line = l.line 
	tok.Column = l.column 

	switch l.ch {
		
	case '{':
		tok = Token{
			Type : LEFT_BRACE,
			Literal: string(l.ch),
			Line: l.line,
			Column: l.column,
		}
	case '}':
		tok = Token{
			Type: RIGHT_BRACE,
			Literal: string(l.ch),
			Line: l.line,
			Column: l.column,
		}
	case 0:
		tok = Token{
			Type: EOF,
			Literal: "",
			Line: l.line,
			Column: l.column,
		}
	case ':':
		tok = Token{
			Type: COLON,
			Literal: ":",
			Line: l.line,
			Column: l.column,
		}
	case ',':
		tok = Token{
			Type: COMMA,
			Literal: ",",
			Line: l.line,
			Column: l.column,
		}
	case '"':
		tok = Token{
			Type: STRING,
			Literal: '"',
			Line: l.line,
			Column: l.column,
		}

	default:
		tok = Token{
			Type: ILLEGAL,
			Literal: string(l.ch),
			Line: l.line,
			Column: l.column,
		}
	
	}
	
	// read the next character 
	l.readChar()
	
	return tok
}
