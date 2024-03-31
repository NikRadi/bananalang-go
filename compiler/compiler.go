package compiler

import (
	"bananalang/ast"
	"bananalang/instruction"
	"bananalang/token"
	"fmt"
	"strconv"
)

func Compile(expression ast.Expression) []instruction.Instruction {
	var instructions []instruction.Instruction
	switch expr := expression.(type) {
	case ast.Literal:
		value, err := strconv.Atoi(expr.Value)
		if err != nil {
			fmt.Println("Compile error: Invalid number")
			return nil
		}

		instructions = append(instructions, instruction.Push)
		instructions = append(instructions, instruction.Instruction(value))
	case ast.BinaryOperator:
		instructions = append(instructions, Compile(expr.LeftExpression)...)
		instructions = append(instructions, Compile(expr.RightExpression)...)
		switch expr.Operator {
		case token.Plus:
			instructions = append(instructions, instruction.Add)
		case token.Minus:
			instructions = append(instructions, instruction.Sub)
		case token.Star:
			instructions = append(instructions, instruction.Mul)
		default:
			fmt.Println("Compile error: Unknown binary operator")
		}
	default:
		fmt.Println("Compile error: Unknown expression type")
	}

	return instructions
}
