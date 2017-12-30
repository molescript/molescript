package lexer

type Lexer struct {
	input    string
	start    int
	position int
	current  byte
	line     int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var token Token

	l.skipWhitespace()

	switch l.current {
	case '(':
		token = newToken(LEFT_PAREN, l.current)
	case ')':
		token = newToken(RIGHT_PAREN, l.current)
	case '{':
		token = newToken(LEFT_BRACE, l.current)
	case '}':
		token = newToken(RIGHT_BRACE, l.current)
	case ',':
		token = newToken(COMMA, l.current)
	case '.':
		token = newToken(DOT, l.current)
	case '+':
		token = newToken(PLUS, l.current)
	case '-':
		token = newToken(MINUS, l.current)
	case ';':
		token = newToken(SEMICOLON, l.current)
	case '*':
		token = newToken(ASTERISK, l.current)
	case '/':
		token = newToken(SLASH, l.current)
	case '=':
		if l.peekChar() == '=' {
			current := l.current
			l.readChar()
			token = Token{Type: EQUAL, Lexeme: string(current) + string(l.current)}
		} else {
			token = newToken(ASSIGN, l.current)
		}
	case '!':
		if l.peekChar() == '=' {
			current := l.current
			l.readChar()
			token = Token{Type: BANG_EQUAL, Lexeme: string(current) + string(l.current)}
		} else {
			token = newToken(BANG, l.current)
		}
	case '>':
		if l.peekChar() == '=' {
			current := l.current
			l.readChar()
			token = Token{Type: GREATER_EQUAL, Lexeme: string(current) + string(l.current)}
		} else {
			token = newToken(GREATER, l.current)
		}
	case '<':
		if l.peekChar() == '=' {
			current := l.current
			l.readChar()
			token = Token{Type: LESS_EQUAL, Lexeme: string(current) + string(l.current)}
		} else {
			token = newToken(LESS, l.current)
		}
	case 0:
		token.Type = EOF
		token.Lexeme = ""
	default:
		if isLetter(l.current) {
			lexeme := l.readWord()
			if keyword, ok := keywords[lexeme]; ok {
				token.Type = keyword
			} else {
				token.Type = IDENTIFIER
			}
			token.Lexeme = lexeme
		} else if isDigit(l.current) {
			number := l.readNumber()
			token.Type = NUMBER
			token.Lexeme = number
		} else {
			token = newToken(ILLEGAL, l.current)
		}
	}
	l.readChar()
	return token
}

func newToken(tokenType TokenType, current byte) Token {
	return Token{Type: tokenType, Lexeme: string(current)}

}

func (l *Lexer) readChar() {
	if l.position >= len(l.input) {
		l.current = 0
	} else {
		l.current = l.input[l.position]
	}

	l.start = l.position
	l.position += 1
}

func (l *Lexer) peekChar() byte {
	if l.position >= len(l.input) {
		return 0
	} else {
		return l.input[l.position]
	}
}

func (l *Lexer) readWord() string {
	start := l.start
	for isLetter(l.peekChar()) || isDigit(l.peekChar()) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.start
	for isDigit(l.peekChar()) || l.peekChar() == '.' {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.current == ' ' || l.current == '\t' || l.current == '\n' || l.current == '\r' {
		l.readChar()
	}
}

func isLetter(b byte) bool {
	return b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' || b == '_'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
