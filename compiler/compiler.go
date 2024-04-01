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
	var instructions []opcode.Opcode
	switch expr := expression.(type) {
	case ast.Literal:
		switch expr.Type {
		case token.LiteralNumber:
			value, err := strconv.Atoi(expr.Value)
			if err != nil {
				fmt.Println("Compile error: Invalid number")
				return nil
			}

			instructions = append(instructions, opcode.Push)
			instructions = append(instructions, opcode.Opcode(value))
		default:
			fmt.Println("Compile error: Invalid literal", expr.Type)
		}
	case ast.BinaryOperator:
		instructions = append(instructions, compiler.Compile(expr.LeftExpression)...)
		instructions = append(instructions, compiler.Compile(expr.RightExpression)...)
		switch expr.Operator {
		case token.Plus:
			instructions = append(instructions, opcode.Add)
		case token.Minus:
			instructions = append(instructions, opcode.Sub)
		case token.Star:
			instructions = append(instructions, opcode.Mul)
		default:
			fmt.Println("Compile error: Unknown binary operator")
		}
	default:
		fmt.Println("Compile error: Unknown expression type")
	}

	return instructions
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
