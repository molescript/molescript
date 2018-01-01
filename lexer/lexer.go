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
	l.line = 1
	return l
}

func (l *Lexer) ReadToken() Token {
	var token Token

	l.skipIneffective()

	switch l.current {
	case '(':
		token = newToken(LEFT_PAREN, l.current, l.line)
	case ')':
		token = newToken(RIGHT_PAREN, l.current, l.line)
	case '{':
		token = newToken(LEFT_BRACE, l.current, l.line)
	case '}':
		token = newToken(RIGHT_BRACE, l.current, l.line)
	case ',':
		token = newToken(COMMA, l.current, l.line)
	case '.':
		token = newToken(DOT, l.current, l.line)
	case '+':
		token = newToken(PLUS, l.current, l.line)
	case '-':
		token = newToken(MINUS, l.current, l.line)
	case ';':
		token = newToken(SEMICOLON, l.current, l.line)
	case '*':
		token = newToken(ASTERISK, l.current, l.line)
	case '/':
		token = newToken(SLASH, l.current, l.line)
	case '=':
		if l.peekChar() == '=' {
			current := l.current
			l.readChar()
			token = Token{Type: EQUAL, Literal: string(current) + string(l.current), Line: l.line}
		} else {
			token = newToken(ASSIGN, l.current, l.line)
		}
	case '!':
		if l.peekChar() == '=' {
			current := l.current
			l.readChar()
			token = Token{Type: BANG_EQUAL, Literal: string(current) + string(l.current), Line: l.line}
		} else {
			token = newToken(BANG, l.current, l.line)
		}
	case '>':
		if l.peekChar() == '=' {
			current := l.current
			l.readChar()
			token = Token{Type: GREATER_EQUAL, Literal: string(current) + string(l.current), Line: l.line}
		} else {
			token = newToken(GREATER, l.current, l.line)
		}
	case '<':
		if l.peekChar() == '=' {
			current := l.current
			l.readChar()
			token = Token{Type: LESS_EQUAL, Literal: string(current) + string(l.current), Line: l.line}
		} else {
			token = newToken(LESS, l.current, l.line)
		}
	case 0:
		token.Type = EOF
		token.Literal = ""
		token.Line = l.line - 1
	default:
		if isLetter(l.current) {
			lexeme := l.readWord()
			if keyword, ok := keywords[lexeme]; ok {
				token.Type = keyword
			} else {
				token.Type = IDENTIFIER
			}
			token.Literal = lexeme
			token.Line = l.line
		} else if isDigit(l.current) {
			number := l.readNumber()
			token.Type = NUMBER
			token.Literal = number
			token.Line = l.line
		} else {
			token = newToken(ILLEGAL, l.current, l.line)
		}
	}
	l.readChar()
	return token
}

func newToken(tokenType TokenType, current byte, line int) Token {
	return Token{Type: tokenType, Literal: string(current), Line: line}

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

func (l *Lexer) skipComments() {
	for l.current != '\n' {
		l.readChar()
	}
	l.readChar()
	l.line += 1
}

func (l *Lexer) skipIneffective() {
	for {
		if l.current == '\n' {
			l.line += 1
		}

		if l.current == ' ' || l.current == '\t' || l.current == '\n' || l.current == '\r' {
			l.readChar()
		} else if l.current == '#' {
			l.skipComments()
		} else {
			break
		}
	}
}

func isLetter(b byte) bool {
	return b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' || b == '_'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
