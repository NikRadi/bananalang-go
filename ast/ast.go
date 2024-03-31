package ast

import (
	"bananalang/token"
)

type (
	Expression interface {}

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
