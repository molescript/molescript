package lexer

type TokenType string

const (
	// Single character tokens
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"
	COMMA       = ","
	DOT         = "."
	PLUS        = "+"
	MINUS       = "-"
	SEMICOLON   = ";"
	ASTERISK    = "*"
	SLASH       = "/"

	// Single or double character tokens
	BANG          = "!"
	BANG_EQUAL    = "!="
	ASSIGN        = "="
	EQUAL         = "=="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="

	// Literals
	IDENTIFIER = "IDENTIFIER"
	STRING     = "STRING"
	NUMBER     = "NUMBER"

	// Keywords
	AND      = "AND"
	CLASS    = "CLASS"
	ELSE     = "ELSE"
	FALSE    = "FALSE"
	FUNCTION = "FUNCTION"
	FOR      = "FOR"
	IF       = "IF"
	NONE     = "NONE"
	OR       = "OR"
	RETURN   = "RETURN"
	SUPER    = "SUPER"
	THIS     = "THIS"
	TRUE     = "TRUE"
	VAR      = "VAR"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"False":  FALSE,
	"func":   FUNCTION,
	"for":    FOR,
	"if":     IF,
	"None":   NONE,
	"or":     OR,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"True":   TRUE,
	"var":    VAR,
}

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}
