package opcode

import (
	"fmt"
)

type Opcode byte

const (
	// Push a value to the stack
	Push Opcode = iota

	// Print the top value of the stack
	Print

	// Pop the top two values, add them, and push the result
	Add

	// Pop the top two values, subtract the 2nd from the 1st, and push the result
	Sub

	// Pop the top two values, multiply them, and push the result
	Mul
)

var opcodes = [...]string{
	Push:	"Push",
	Print:	"Print",
	Add:	"Add",
	Sub:	"Sub",
	Mul:	"Mul",
}

func (opcode Opcode) String() string {
	return opcodes[opcode]
}

func PrintOpcodes(instructions []Opcode) {
	for i := 0; i < len(instructions); i += 1 {
		instr := instructions[i]
		switch instr {
		case Push:
			i += 1
			fmt.Printf("%-10s %d\n", instr, instructions[i])
		default:
			fmt.Println(instr)
		}
	}
}
