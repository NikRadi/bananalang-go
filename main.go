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
	const code = "a=1; b=2; c=3; d=4; a+b+c+d"
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
