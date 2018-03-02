package parser

import (
	"bytes"
	"molescript/lexer"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
	Errors     []ParseError
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, smt := range p.Statements {
		out.WriteString(smt.String())
	}

	return out.String()
}

type IdentifierExpr struct {
	Token lexer.Token
	Value string
}

func (i *IdentifierExpr) expressionNode()      {}
func (i *IdentifierExpr) TokenLiteral() string { return i.Token.Literal }
func (i *IdentifierExpr) String() string       { return i.Token.Literal }

type AssignStmt struct {
	Token lexer.Token
	Name  *IdentifierExpr
	Value Expression
}

func (s *AssignStmt) statementNode()       {}
func (s *AssignStmt) TokenLiteral() string { return s.Token.Literal }
func (s *AssignStmt) String() string {
	var out bytes.Buffer

	out.WriteString(s.Name.String() + " = ")
	if s.Value != nil {
		out.WriteString(s.Value.String())
	}

	return out.String()
}

type ReturnStmt struct {
	Token lexer.Token
	Value Expression
}

func (s *ReturnStmt) statementNode()       {}
func (s *ReturnStmt) TokenLiteral() string { return s.Token.Literal }
func (s *ReturnStmt) String() string {
	var out bytes.Buffer

	out.WriteString("return ")
	if s.Value != nil {
		out.WriteString(s.Value.String())
	}

	return out.String()
}
