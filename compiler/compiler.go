package compiler

import (
	"bananalang/ast"
	"bananalang/opcode"
	"bananalang/token"
	"fmt"
	"strconv"
)

type (
	symbol struct {
		identifier 	string
		offset		int
	}

	Compiler struct {
		instructions 	[]opcode.Opcode
	}
)

func NewCompiler() *Compiler {
	return &Compiler{
		instructions:	[]opcode.Opcode{},
	}
}

func (compiler *Compiler) Compile(expression ast.Expression) []opcode.Opcode {
	switch expr := expression.(type) {
	case ast.Literal:
		switch expr.Type {
		case token.LiteralNumber:
			value, _ := strconv.Atoi(expr.Value)
			compiler.emit(opcode.Push, opcode.Opcode(value))
		default:
			fmt.Println("Compile error: unknown literal")
		}
	case ast.BinaryOperator:
		compiler.Compile(expr.LeftExpression)
		compiler.Compile(expr.RightExpression)
		switch expr.Operator {
		case token.Plus:
			compiler.emit(opcode.Add)
		case token.Minus:
			compiler.emit(opcode.Sub)
		case token.Star:
			compiler.emit(opcode.Mul)
		default:
			fmt.Println("Compile error: unknown binary operator")
		}
	default:
		fmt.Printf("Compile error: unknown expression type: %T\n", expr)
	}

	return compiler.instructions
}

func (compiler *Compiler) emit(codes ...opcode.Opcode) {
	compiler.instructions = append(compiler.instructions, codes...)
}
