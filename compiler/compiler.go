package compiler

import (
	"bananalang/ast"
	"bananalang/opcode"
	"bananalang/token"
	"fmt"
	"os"
	"strconv"
)

type Compiler struct {
	instructions 	[]opcode.Opcode
}

func NewCompiler() *Compiler {
	return &Compiler{
		instructions:	[]opcode.Opcode{},
	}
}

func (compiler *Compiler) Compile(statement ast.ExpressionStatement) []opcode.Opcode {
	for _, expression := range statement.Expressions {
		compiler.compileExpression(expression)
	}

	return compiler.instructions
}

func (compiler *Compiler) compileExpression(expression ast.Expression) {
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
		compiler.compileExpression(expr.LeftExpression)
		compiler.compileExpression(expr.RightExpression)
		switch expr.Operator {
		case token.Plus:
			compiler.emit(opcode.Add)
		case token.Minus:
			compiler.emit(opcode.Sub)
		case token.Star:
			compiler.emit(opcode.Mul)
		case token.TwoEquals:
			compiler.emit(opcode.CmpEqu)
		case token.NotEquals:
			compiler.emit(opcode.CmpNeq)
		case token.LessThan:
			compiler.emit(opcode.CmpLet)
		case token.LessThanEquals:
			compiler.emit(opcode.CmpLte)
		case token.GreaterThan:
			compiler.emit(opcode.CmpGrt)
		case token.GreaterThanEquals:
			compiler.emit(opcode.CmpGte)
		default:
			fmt.Println("Compile error: unknown binary operator")
			os.Exit(1)
		}
	case ast.UnaryOperator:
		compiler.compileExpression(expr.Expression)
		switch expr.Operator {
		case token.Minus:
			compiler.emit(opcode.Neg)
		default:
			fmt.Println("Compile error: unknown unary operator")
			os.Exit(1)
		}
	default:
		fmt.Printf("Compile error: unknown expression type: %T\n", expr)
		os.Exit(1)
	}
}

func (compiler *Compiler) emit(codes ...opcode.Opcode) {
	compiler.instructions = append(compiler.instructions, codes...)
}
