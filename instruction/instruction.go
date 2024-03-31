package instruction

import (
	"fmt"
)

type Instruction byte

const (
	// Push a value to the stack
	Push Instruction = iota

	// Print the top value of the stack
	Print

	// Pop the top two values, add them, and push the result
	Add

	// Pop the top two values, subtract the 2nd from the 1st, and push the result
	Sub

	// Pop the top two values, multiply them, and push the result
	Mul
)

func PrintInstructions(instructions []Instruction) {
	for i := 0; i < len(instructions); i += 1 {
		instruction := instructions[i]
		switch instruction {
		case Push:
			i += 1
			value := instructions[i]
			fmt.Println("Push", value)
		case Print:
			fmt.Println("Print")
		case Add:
			fmt.Println("Add")
		case Sub:
			fmt.Println("Sub")
		case Mul:
			fmt.Println("Mul")
		}
	}
}
