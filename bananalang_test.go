package main

import (
	"bananalang/compiler"
	"bananalang/lexer"
	"bananalang/parser"
	"bananalang/vm"
	"testing"
)

func executeAndExpect(t *testing.T, code string, expected int) {
	lex := lexer.NewLexer(code)
	par := parser.NewParser(lex)
	tree := par.Parse()

	com := compiler.NewCompiler()
	instructions := com.Compile(tree)

	runtime := vm.NewVM()
	runtime.Execute(instructions)

	frame := runtime.Stack[0]
	index := len(frame.OperandStack) - 1
	actual := frame.OperandStack[index]
	if actual != expected {
		t.Errorf("Expected %d but got %d (%s)", expected, actual, code)
	}
}

func TestArithmetic(t *testing.T) {
	executeAndExpect(t, "1+2;", 3)
	executeAndExpect(t, "3-2;", 1)
	executeAndExpect(t, "2   -    2;", 0)
	executeAndExpect(t, "2 + 3 * 2;", 8)
	executeAndExpect(t, "2 - 3 * 2;", -4)
	executeAndExpect(t, "2 * 3 + 2;", 8)
	executeAndExpect(t, "2 * 3 - 2;", 4)
	executeAndExpect(t, "2 + (3 - 3);", 2)
	executeAndExpect(t, "(2 * 3) - 2;", 4)
	executeAndExpect(t, "2 * (1 + 1);", 4)
	executeAndExpect(t, "-1;", -1)
	executeAndExpect(t, "-(1 + 1);", -2)
	executeAndExpect(t, "-3 * -2;", 6)
	executeAndExpect(t, "-4 + -4;", -8)
	executeAndExpect(t, "-----7;", -7)
	executeAndExpect(t, "1 + ++7;", 8)

	executeAndExpect(t, "1==1;", 1)
	executeAndExpect(t, "2==1;", 0)
	executeAndExpect(t, "1!=1;", 0)
	executeAndExpect(t, "2!=1;", 1)

	executeAndExpect(t, "1 > 1;", 0)
	executeAndExpect(t, "0 > 1;", 0)
	executeAndExpect(t, "2 > 1;", 1)

	executeAndExpect(t, "1 >= 1;", 1)
	executeAndExpect(t, "0 >= 1;", 0)
	executeAndExpect(t, "2 >= 1;", 1)

	executeAndExpect(t, "1 < 1;", 0)
	executeAndExpect(t, "0 < 1;", 1)
	executeAndExpect(t, "2 < 1;", 0)

	executeAndExpect(t, "1 <= 1;", 1)
	executeAndExpect(t, "0 <= 1;", 1)
	executeAndExpect(t, "2 <= 1;", 0)

	executeAndExpect(t, "1+1; 2+2; 3+3", 6)

	executeAndExpect(t, "a=2; a", 2)
	executeAndExpect(t, "a=1+2; a=a*3; a;", 9)
}
