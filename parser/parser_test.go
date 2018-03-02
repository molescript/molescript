package parser

import (
	"testing"

	"molescript/lexer"
)

func TestAssignStatements(t *testing.T) {
	input := `x = 5;
y = 10;
z = 15;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d",
			len(program.Statements))
	}

	stmtTests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"z"},
	}

	for i, tt := range stmtTests {
		stmt := program.Statements[i]
		if !testAssignStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testAssignStatement(t *testing.T, s Statement, name string) bool {
	if s.TokenLiteral() != name {
		t.Errorf("s.TokenLiteral not '%s'. got=%q", name, s.TokenLiteral())
		return false
	}

	assignStmt, ok := s.(*AssignStmt)
	if !ok {
		t.Errorf("s not AssignStmt. got=%T", s)
		return false
	}

	if assignStmt.Name.Value != name {
		t.Errorf("assignStmt.Name.Expn not '%s'. got=%s", name, assignStmt.Name.Value)
		return false
	}

	if assignStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, assignStmt.Name)
		return false
	}

	return true
}

func TestParseErrors(t *testing.T) {
	input := `x = 5;
foo 42;
bar 1337;
y = 10;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d",
			len(program.Statements))
	}

	if len(program.Errors) != 2 {
		t.Fatalf("program.Errors does not contain 2 erros. got=%d",
			len(program.Errors))
	}

	errTests := []struct {
		line     int
		tokType  lexer.TokenType
		literal  string
		expected lexer.TokenType
	}{
		{2, lexer.NUMBER, "42", lexer.ASSIGN},
		{3, lexer.NUMBER, "1337", lexer.ASSIGN},
	}

	for i, tt := range errTests {
		err := program.Errors[i]
		if !testParseError(t, err, tt.line, tt.tokType, tt.literal, tt.expected) {
			return
		}
	}

}

func testParseError(t *testing.T, e ParseError, line int, tokType lexer.TokenType, literal string, expected lexer.TokenType) bool {
	if e.tok.Line != line {
		t.Errorf("s.Line number not %d. got=%d", line, e.tok.Line)
		return false
	}

	if e.tok.Type != tokType {
		t.Errorf("s.Type not '%s'. got=%q", tokType, e.tok.Type)
		return false
	}

	if e.tok.Literal != literal {
		t.Errorf("s.Literal not '%s'. got=%q", literal, e.tok.Literal)
		return false
	}

	if e.expected != expected {
		t.Errorf("s.Literal not '%s'. got=%q", expected, e.expected)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `return 2;
return 5;
return 3.14;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	stmtTests := []struct {
		expectedValue string
	}{
		{"2"},
		{"5"},
		{"3.14"},
	}

	for i, tt := range stmtTests {
		stmt := program.Statements[i]
		if !testReturnStatement(t, stmt, tt.expectedValue) {
			return
		}
	}
}

func testReturnStatement(t *testing.T, s Statement, value string) bool {
	returnStmt, ok := s.(*ReturnStmt)
	if !ok {
		t.Errorf("s not ReturnStmt. got=%T", s)
		return false
	}

	if returnStmt.TokenLiteral() != "return" {
		t.Errorf("s.Name not '%s'. got=%s", value, returnStmt.Value)
		return false
	}

	return true
}
