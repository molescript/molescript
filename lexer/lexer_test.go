package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `five = 5;
pi = 3.14159265359; # This is a comment
# This too is a comment

add = func(x, y) {
  x + y;
};

result = add(five, ten);
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
		line            int
		expectedType    TokenType
		expectedLiteral string
	}{
		{1, IDENTIFIER, "five"},
		{1, ASSIGN, "="},
		{1, NUMBER, "5"},
		{1, SEMICOLON, ";"},
		{2, IDENTIFIER, "pi"},
		{2, ASSIGN, "="},
		{2, NUMBER, "3.14159265359"},
		{2, SEMICOLON, ";"},
		{5, IDENTIFIER, "add"},
		{5, ASSIGN, "="},
		{5, FUNCTION, "func"},
		{5, LEFT_PAREN, "("},
		{5, IDENTIFIER, "x"},
		{5, COMMA, ","},
		{5, IDENTIFIER, "y"},
		{5, RIGHT_PAREN, ")"},
		{5, LEFT_BRACE, "{"},
		{6, IDENTIFIER, "x"},
		{6, PLUS, "+"},
		{6, IDENTIFIER, "y"},
		{6, SEMICOLON, ";"},
		{7, RIGHT_BRACE, "}"},
		{7, SEMICOLON, ";"},
		{9, IDENTIFIER, "result"},
		{9, ASSIGN, "="},
		{9, IDENTIFIER, "add"},
		{9, LEFT_PAREN, "("},
		{9, IDENTIFIER, "five"},
		{9, COMMA, ","},
		{9, IDENTIFIER, "ten"},
		{9, RIGHT_PAREN, ")"},
		{9, SEMICOLON, ";"},
		{10, BANG, "!"},
		{10, MINUS, "-"},
		{10, SLASH, "/"},
		{10, ASTERISK, "*"},
		{10, NUMBER, "5"},
		{10, SEMICOLON, ";"},
		{11, NUMBER, "5"},
		{11, LESS, "<"},
		{11, NUMBER, "10"},
		{11, GREATER, ">"},
		{11, NUMBER, "5"},
		{11, SEMICOLON, ";"},
		{13, IF, "if"},
		{13, LEFT_PAREN, "("},
		{13, NUMBER, "5"},
		{13, LESS, "<"},
		{13, NUMBER, "10"},
		{13, RIGHT_PAREN, ")"},
		{13, LEFT_BRACE, "{"},
		{14, RETURN, "return"},
		{14, TRUE, "True"},
		{14, SEMICOLON, ";"},
		{15, RIGHT_BRACE, "}"},
		{15, ELSE, "else"},
		{15, LEFT_BRACE, "{"},
		{16, RETURN, "return"},
		{16, FALSE, "False"},
		{16, SEMICOLON, ";"},
		{17, RIGHT_BRACE, "}"},
		{19, NUMBER, "10"},
		{19, EQUAL, "=="},
		{19, NUMBER, "10"},
		{19, SEMICOLON, ";"},
		{20, NUMBER, "10"},
		{20, BANG_EQUAL, "!="},
		{20, NUMBER, "9"},
		{20, SEMICOLON, ";"},
		{21, NUMBER, "10"},
		{21, GREATER_EQUAL, ">="},
		{21, NUMBER, "9"},
		{21, SEMICOLON, ";"},
		{22, NUMBER, "9"},
		{22, LESS_EQUAL, "<="},
		{22, NUMBER, "10"},
		{22, SEMICOLON, ";"},
		{24, IDENTIFIER, "a1"},
		{24, DOT, "."},
		{24, IDENTIFIER, "b"},
		{24, ASSIGN, "="},
		{24, IDENTIFIER, "c2"},
		{24, SEMICOLON, ";"},
		{25, ILLEGAL, "@"},
		{25, EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.ReadToken()
		if tok.Line != tt.line {
			t.Fatalf("tests[%d] - line number wrong. expected=%d, got=%d",
				i, tt.line, tok.Line)
		}

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
