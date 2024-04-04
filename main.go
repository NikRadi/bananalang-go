package main

import (
	"bananalang/compiler"
	"bananalang/lexer"
	"bananalang/opcode"
	"bananalang/parser"
	"bananalang/vm"
	"fmt"
)

func main() {
	const code = "abcd = 1; abcd;"
	fmt.Println(code)
	lex := lexer.NewLexer(code)
	par := parser.NewParser(lex)

	tree := par.Parse()

	com := compiler.NewCompiler()
	instructions := com.Compile(tree)
	instructions = append(instructions, opcode.Print)
	opcode.PrintOpcodes(instructions)

	runtime := vm.NewVM()
	runtime.Execute(instructions)
	fmt.Println("=====")
	fmt.Printf("Bytecodes: %d\n", len(instructions))
}
