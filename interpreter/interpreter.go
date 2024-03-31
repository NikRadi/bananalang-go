package interpreter

import (
	"bananalang/instruction"
	"fmt"
)

func Interpret(instructions []instruction.Instruction) {
	var stack [64]int
	sp := 0
	for i := 0; i < len(instructions); i += 1 {
		instr := instructions[i]
		switch instr {
		case instruction.Push:
			i += 1
			value := instructions[i]
			stack[sp] = int(value)
			sp += 1
		case instruction.Print:
			top := stack[sp - 1]
			fmt.Println(top)
		case instruction.Add:
			value1 := stack[sp - 2]
			value2 := stack[sp - 1]
			stack[sp - 2] = value1 + value2
			sp -= 1
		case instruction.Sub:
			value1 := stack[sp - 2]
			value2 := stack[sp - 1]
			stack[sp - 2] = value1 - value2
			sp -= 1
		case instruction.Mul:
			value1 := stack[sp - 2]
			value2 := stack[sp - 1]
			stack[sp - 2] = value1 * value2
			sp -= 1
		}
	}
}
