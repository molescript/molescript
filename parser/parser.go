package parser

import (
	"errors"
	"fmt"
	"molescript/lexer"
)

type ParseError struct {
	tok      lexer.Token
	expected lexer.TokenType
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("Line %d: Expected %s, got %s (%s)", p.tok.Line, p.expected, p.tok.Type, p.tok.Literal)
}

type Parser struct {
	lex     *lexer.Lexer
	current lexer.Token
	next    lexer.Token
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex}

	p.advanceToken()
	p.advanceToken()

	return p
}

func (p *Parser) advanceToken() {
	p.current = p.next
	p.next = p.lex.ReadToken()
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}
	program.Errors = []ParseError{}

	for p.current.Type != lexer.EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			if parseError, ok := err.(*ParseError); ok {
				program.Errors = append(program.Errors, *parseError)
			}

		} else {
			program.Statements = append(program.Statements, stmt)
		}

		p.advanceToken()
	}

	return program
}

func (p *Parser) parseStatement() (Statement, error) {
	switch p.current.Type {
	case lexer.IDENTIFIER:
		return p.parseAssignStatement()
	default:
		msg := fmt.Sprintf("parseStatement() failed, no rule for %s (%s)", p.current.Type, p.current.Literal)
		return nil, errors.New(msg)

	}
}

func (p *Parser) parseAssignStatement() (*AssignStmt, error) {
	stmt := &AssignStmt{Token: p.current}

	stmt.Name = &IdentifierExpr{Token: p.current, Value: p.current.Literal}

	if err := p.expectNext(lexer.ASSIGN); err != nil {
		return nil, &ParseError{p.next, lexer.ASSIGN}
	}

	for p.current.Type != lexer.SEMICOLON {
		p.advanceToken()
	}

	return stmt, nil

}

func (p *Parser) expectNext(tokenType lexer.TokenType) error {
	if p.next.Type == tokenType {
		p.advanceToken()
		return nil
	} else {
		return &ParseError{p.next, tokenType}
	}
}
