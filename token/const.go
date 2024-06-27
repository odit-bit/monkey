package token

const (
	// // comment
	// COMMENT = "//"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals

	IDENT = "IDENT" // add, foobar, x, y, ...

	// DATA types
	
	INT    = "INT" // 1343456
	STRING = "STRING"

	// Operators

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	EQ     = "=="
	NOT_EQ = "!="

	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	// Delimiters

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
