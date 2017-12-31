package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `var five = 5;
var pi = 3.14159265359; # This is a comment
# This too is a comment

var add = func(x, y) {
  x + y;
};

var result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return True;
} else {
	return False;
}

10 == 10;
10 != 9;
10 >= 9;
9 <= 10;

a1.b = c2;
@
`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{VAR, "var"},
		{IDENTIFIER, "five"},
		{ASSIGN, "="},
		{NUMBER, "5"},
		{SEMICOLON, ";"},
		{VAR, "var"},
		{IDENTIFIER, "pi"},
		{ASSIGN, "="},
		{NUMBER, "3.14159265359"},
		{SEMICOLON, ";"},
		{VAR, "var"},
		{IDENTIFIER, "add"},
		{ASSIGN, "="},
		{FUNCTION, "func"},
		{LEFT_PAREN, "("},
		{IDENTIFIER, "x"},
		{COMMA, ","},
		{IDENTIFIER, "y"},
		{RIGHT_PAREN, ")"},
		{LEFT_BRACE, "{"},
		{IDENTIFIER, "x"},
		{PLUS, "+"},
		{IDENTIFIER, "y"},
		{SEMICOLON, ";"},
		{RIGHT_BRACE, "}"},
		{SEMICOLON, ";"},
		{VAR, "var"},
		{IDENTIFIER, "result"},
		{ASSIGN, "="},
		{IDENTIFIER, "add"},
		{LEFT_PAREN, "("},
		{IDENTIFIER, "five"},
		{COMMA, ","},
		{IDENTIFIER, "ten"},
		{RIGHT_PAREN, ")"},
		{SEMICOLON, ";"},
		{BANG, "!"},
		{MINUS, "-"},
		{SLASH, "/"},
		{ASTERISK, "*"},
		{NUMBER, "5"},
		{SEMICOLON, ";"},
		{NUMBER, "5"},
		{LESS, "<"},
		{NUMBER, "10"},
		{GREATER, ">"},
		{NUMBER, "5"},
		{SEMICOLON, ";"},
		{IF, "if"},
		{LEFT_PAREN, "("},
		{NUMBER, "5"},
		{LESS, "<"},
		{NUMBER, "10"},
		{RIGHT_PAREN, ")"},
		{LEFT_BRACE, "{"},
		{RETURN, "return"},
		{TRUE, "True"},
		{SEMICOLON, ";"},
		{RIGHT_BRACE, "}"},
		{ELSE, "else"},
		{LEFT_BRACE, "{"},
		{RETURN, "return"},
		{FALSE, "False"},
		{SEMICOLON, ";"},
		{RIGHT_BRACE, "}"},
		{NUMBER, "10"},
		{EQUAL, "=="},
		{NUMBER, "10"},
		{SEMICOLON, ";"},
		{NUMBER, "10"},
		{BANG_EQUAL, "!="},
		{NUMBER, "9"},
		{SEMICOLON, ";"},
		{NUMBER, "10"},
		{GREATER_EQUAL, ">="},
		{NUMBER, "9"},
		{SEMICOLON, ";"},
		{NUMBER, "9"},
		{LESS_EQUAL, "<="},
		{NUMBER, "10"},
		{SEMICOLON, ";"},
		{IDENTIFIER, "a1"},
		{DOT, "."},
		{IDENTIFIER, "b"},
		{ASSIGN, "="},
		{IDENTIFIER, "c2"},
		{SEMICOLON, ";"},
		{ILLEGAL, "@"},
		{EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.ReadToken()
		//fmt.Println(tests[i], tok.Lexeme)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
