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
)

func (Literal)			expression() {}
func (BinaryOperator)	expression() {}
