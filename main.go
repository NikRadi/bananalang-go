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
	const code = "a=2; a=a*3"
	lex := lexer.NewLexer(code)
	par := parser.NewParser(lex)

	tree := par.Parse()
	fmt.Println(tree)

	com := compiler.NewCompiler()
	instructions := com.Compile(tree)
	opcode.PrintOpcodes(instructions)

	runtime := vm.NewVM()
	runtime.Execute(instructions)
	fmt.Println("=====")
	fmt.Printf("Bytecodes: %d\n", len(instructions))
}
