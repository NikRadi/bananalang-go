package ast

import (
	"bananalang/token"
)

type Expression interface {
	expression()
}

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
