package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" 	// undefined tokens are illegal
	EOF		= "EOF"			// end of file

	// Identifiers + literals
	IDENT 	= "IDENT"
	INT 	= "INT"

	// Operators
	ASSIGN 	 = "="
	PLUS 	 = "+"
	MINUS 	 = "-"
	BANG 	 = "!"
	ASTERISK = "*"
	SLASH	 = "/"

	LT = "<"
	GT = ">"

	// Delimiters
	COMMA	  = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET 	 = "LET"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

/*
Checks the keywords table to see whether the given identifier is in fact a keyword.
If it is, it returns the keyword's TokenType constant, otherwise we get token.IDENT to
signify a user-defined identifier.
*/
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}