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
		symbols			map[string]int // Variable name to stack offset
	}
)

func NewCompiler() *Compiler {
	return &Compiler{
		instructions:	[]opcode.Opcode{},
		symbols:		make(map[string]int),
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

func (compiler *Compiler) addSymbol(name string) {
	if _, exists := compiler.symbols[name]; exists {
		fmt.Println("Compile error: Symbol already exists", name)
	}

	compiler.symbols[name] = len(compiler.symbols)
}

func (compiler *Compiler) getSymbolOffset(name string) int {
	if offset, exists := compiler.symbols[name]; exists {
		return offset
	}

	return -1
}
