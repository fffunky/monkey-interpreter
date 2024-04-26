package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input		 string
	position	 int	// current position in input (points to current char)
	readPosition int 	// current reading position in input (after current char)
	ch 			 byte 	// current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/*
Advances the lexer by one character and stores the current character in the ch field.
If the end of the file is reached then ch is given a value of 0.
*/
func (l *Lexer) readChar() {
	if (l.readPosition >= len(l.input)) {
		l.ch = 0	// 0 is ascii for "NUL" -> either nothing is read or "end of file"
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

/*
Looks at the current character under examination (l.ch) and returns a token depending on which character it is.
Advances the input pointers before returning the token so l.ch is already updated for the next call.
*/
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch (l.ch) {
		case '=':
			tok = newToken(token.ASSIGN, l.ch)
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}':
			tok = newToken(token.RBRACE, l.ch)
		case ',':
			tok = newToken(token.COMMA, l.ch)
		case '+':
			tok = newToken(token.PLUS, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			if isLetter(l.ch) {
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				return tok
			} else if isDigit(l.ch) {
				tok.Type = token.INT
				tok.Literal = l.readNumber()
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.ch)
			}
	}

	l.readChar()
	return tok
}

/*
Looks at the current character under examination (l.ch) and advances the
lexer's position if l.ch is a whitespace character.
*/
func (l *Lexer) skipWhitespace() {
	for (l.ch == ' ') || (l.ch == '\t') || (l.ch == '\n') || (l.ch == '\r') {
		l.readChar()
	}
}

/*
Reads in a number and advances the lexer's position until it encounters 
a non-digit character.
*/
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/*
Reads in an identifier and advances the lexer's position until it encounters 
a non-letter character.
*/
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/*
Determines whether the passed char is a valid letter. All chars a-z, A-Z,
or '_' are valid letters.
*/
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || (ch == '_')
}

/*
Determines whether the passed char is a valid digit (0-9).
*/
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

/*
Initializes a new token.
*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}