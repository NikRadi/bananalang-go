package main

import (
	"bananalang/compiler"
	"bananalang/lexer"
	"bananalang/parser"
	"bananalang/instruction"
	"bananalang/interpreter"
)

func main() {
	const code = "1   +   2  *3-  1"
	lex := lexer.NewLexer(code)
	par := parser.NewParser(lex)
	tree := par.Parse()
	instructions := compiler.Compile(tree)
	instructions = append(instructions, instruction.Print)
	instruction.PrintInstructions(instructions)
	interpreter.Interpret(instructions)
}
