package opcode

import (
	"fmt"
)

type Opcode byte

const (
	// Push a value to the stack
	Push Opcode = iota

	// Pop the top value of the stack
	Pop

	// Print the top value of the stack
	Print

	// Pop the top two values, push 2nd value + 1st value
	Add

	// Pop the top two values, push 2nd value - 1st value
	Sub

	// Pop the top two values, push 2nd value * 1st value
	Mul

	// Pop the top value, push -value
	Neg

	// Pop the two top values, if 2nd value == 1st value push 1, otherwise push 0
	CmpEqu

	// Pop the two top values, if 2nd value != 1st value push 1, otherwise push 0
	CmpNeq

	// Pop the two top values, if 2nd value < 1st value push 1, otherwise push 0
	CmpLet

	// Pop the two top values, if 2nd value <= 1st value push 1, otherwise push 0
	CmpLte

	// Pop the two top values, if 2nd value > 1st value push 1, otherwise push 0
	CmpGrt

	// Pop the two top values, if 2ndl value >= 1st value push 1, otherwise push 0
	CmpGte
)

var opcodes = [...]string{
	Push:   "Push",
	Print:  "Print",
	Add:    "Add",
	Sub:    "Sub",
	Mul:    "Mul",
	Neg:    "Neg",
	CmpEqu: "CmpEqu",
	CmpNeq: "CmpNeq",
	CmpLet: "CmpLet",
	CmpLte: "CmpLte",
	CmpGrt: "CmpGrt",
	CmpGte: "CmpGte",
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
