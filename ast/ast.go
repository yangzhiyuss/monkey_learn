package ast

import "monkey_learn/token"

type Node interface {
	TokenLiteral() string
}

// Statement 所有的Statement(语句)继承这个语句(不产生值)
type Statement interface {
	Node
	statementNode()
}

// Expression 所有的Expression(表达式)继承这个语句(产生值)
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement Statements
type LetStatement struct {
	Token token.Token //token.LET的词法单元
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier Expressions
type Identifier struct {
	Token token.Token //token.IDENT词法单元
	Value string
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
