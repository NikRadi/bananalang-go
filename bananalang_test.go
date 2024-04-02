package main

import (
	"bananalang/compiler"
	"bananalang/lexer"
	"bananalang/parser"
	"bananalang/opcode"
	"bananalang/vm"
	"testing"
)

func executeAndExpect(t *testing.T, code string, expected int) {
	lex := lexer.NewLexer(code)
	par := parser.NewParser(lex)
	tree := par.Parse()

	com := compiler.NewCompiler()
	instructions := com.Compile(tree)
	instructions = append(instructions, opcode.Print)

	runtime := vm.NewVM()
	runtime.Execute(instructions)
	actual := runtime.LastPoppedInt()
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestArithmetic(t *testing.T) {
	executeAndExpect(t, "1+2", 3)
	executeAndExpect(t, "3-2", 1)
	executeAndExpect(t, "2   -    2", 0)
	executeAndExpect(t, "2 + 3 * 2", 8)
	executeAndExpect(t, "2 - 3 * 2", -4)
	executeAndExpect(t, "2 * 3 + 2", 8)
	executeAndExpect(t, "2 * 3 - 2", 4)
}
