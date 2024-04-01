package vm

import (
	"bananalang/opcode"
	"fmt"
)

const stackSize = 8

type VM struct {
	stack 	[stackSize]int
	sp 		int
}

func NewVM() *VM {
	return &VM{
		stack: 	[stackSize]int{},
		sp:		0,
	}
}

func (vm *VM) Execute(codes []opcode.Opcode) {
	for i := 0; i < len(codes); i += 1 {
		code := codes[i]
		switch code {
		case opcode.Push:
			i += 1
			value := int(codes[i])
			vm.push(value)
		case opcode.Print:
			value := vm.pop()
			fmt.Println(value)
		case opcode.Add:
			value1 := vm.pop()
			value2 := vm.pop()
			vm.push(value2 + value1)
		case opcode.Sub:
			value1 := vm.pop()
			value2 := vm.pop()
			vm.push(value2 - value1)
		case opcode.Mul:
			value1 := vm.pop()
			value2 := vm.pop()
			vm.push(value2 * value1)
		default:
			fmt.Println("Runtime error: Unknown instruction", code)
		}
	}
}

func (vm *VM) pop() int {
	vm.sp -= 1
	return vm.stack[vm.sp]
}

func (vm *VM) push(value int) {
	vm.stack[vm.sp] = value
	vm.sp += 1
}
