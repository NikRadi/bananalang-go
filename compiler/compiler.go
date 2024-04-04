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
	instructions         []opcode.Opcode
	variableToStackIndex map[string]int
}

func NewCompiler() *Compiler {
	return &Compiler{
		instructions:         []opcode.Opcode{},
		variableToStackIndex: make(map[string]int),
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
		case token.Identifier:
			index := compiler.variableToStackIndex[expr.Value]
			compiler.emit(opcode.Load, opcode.Opcode(index))
		default:
			fmt.Println("Compile error: unknown literal", expr)
			os.Exit(1)
		}
	case ast.BinaryOperator:
		if expr.Operator == token.Equals {
			compiler.compileExpression(expr.RightExpression)
			index := compiler.getIndex(expr.LeftExpression)
			compiler.emit(opcode.Store, opcode.Opcode(index))
			break
		}

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

func (compiler *Compiler) getIndex(expression ast.Expression) int {
	switch expr := expression.(type) {
	case ast.Literal:
		index, ok := compiler.variableToStackIndex[expr.Value]
		if ok {
			return index
		}

		index = len(compiler.variableToStackIndex)
		compiler.variableToStackIndex[expr.Value] = index
		return index
	default:
		fmt.Println("Compile error: cannot find index of type", expression)
		os.Exit(1)
		return -1
	}
}

func (compiler *Compiler) emit(codes ...opcode.Opcode) {
	compiler.instructions = append(compiler.instructions, codes...)
}
