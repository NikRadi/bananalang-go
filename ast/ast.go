package ast

import (
	"bananalang/token"
)

type (
	Expression interface {
		expression()
	}

	Statement interface {
		statement()
	}
)

// Expressions
type (
	Literal struct {
		Type 	token.Type
		Value 	string
	}

	BinaryOperator struct {
		Operator 		token.Type
		LeftExpression	Expression
		RightExpression Expression
	}

	UnaryOperator struct {
		Operator		token.Type
		Expression 		Expression
	}
)

func (Literal)			expression() {}
func (BinaryOperator)	expression() {}
func (UnaryOperator)	expression() {}

// Statements
type (
	ExpressionStatement struct {
		// TODO: Currently has an Expression slice for development purposes.
		// 		 Should be changed to a single Expression.
		Expressions []Expression
	}
)

func (ExpressionStatement) 	statement() {}
