package parser

import (
	"molescript/lexer"
)

type Parser struct {
	lex *lexer.Lexer

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

	for p.current.Type != lexer.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.advanceToken()
	}

	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.current.Type {
	case lexer.IDENTIFIER:
		return p.parseAssignStatement()
	default:
		return nil

	}
}

func (p *Parser) parseAssignStatement() *AssignStmt {
	stmt := &AssignStmt{Token: p.current}

	stmt.Name = &IdentifierExpr{Token: p.current, Value: p.current.Literal}

	if !p.expectNext(lexer.ASSIGN) {
		return nil
	}

	for p.current.Type != lexer.SEMICOLON {
		p.advanceToken()
	}

	return stmt

}

func (p *Parser) expectNext(tokenType lexer.TokenType) bool {
	if p.next.Type == tokenType {
		p.advanceToken()
		return true
	}

	return false
}
