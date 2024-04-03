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
	const code = "1+2;3+4;"
	lex := lexer.NewLexer(code)
	par := parser.NewParser(lex)

	tree := par.Parse()
	fmt.Println(tree)

	com := compiler.NewCompiler()
	instructions := com.Compile(tree)

	instructions = append(instructions, opcode.Print)
	opcode.PrintOpcodes(instructions)

	runtime := vm.NewVM()
	runtime.Execute(instructions)
	fmt.Println("=====")
	fmt.Printf("Bytecodes: %d\n", len(instructions))
}
