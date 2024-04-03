package vm

import (
	"bananalang/opcode"
	"fmt"
	"os"
)

const stackSize = 8

type VM struct {
	stack [stackSize]int
	sp    int
}

func NewVM() *VM {
	return &VM{
		stack: [stackSize]int{},
		sp:    0,
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
		case opcode.Pop:
			vm.pop()
		case opcode.Print:
			value := vm.pop()
			fmt.Println(value)
		case opcode.Add:
			vm.binaryop(func(v1, v2 int) int { return v2 + v1 })
		case opcode.Sub:
			vm.binaryop(func(v1, v2 int) int { return v2 - v1 })
		case opcode.Mul:
			vm.binaryop(func(v1, v2 int) int { return v2 * v1 })
		case opcode.Neg:
			value := vm.pop()
			vm.push(-value)
		case opcode.CmpEqu:
			vm.binaryop(func(v1, v2 int) int { return boolToInt(v2 == v1) })
		case opcode.CmpNeq:
			vm.binaryop(func(v1, v2 int) int { return boolToInt(v2 != v1) })
		case opcode.CmpLet:
			vm.binaryop(func(v1, v2 int) int { return boolToInt(v2 < v1) })
		case opcode.CmpLte:
			vm.binaryop(func(v1, v2 int) int { return boolToInt(v2 <= v1) })
		case opcode.CmpGrt:
			vm.binaryop(func(v1, v2 int) int { return boolToInt(v2 > v1) })
		case opcode.CmpGte:
			vm.binaryop(func(v1, v2 int) int { return boolToInt(v2 >= v1) })
		default:
			fmt.Println("Runtime error: Unknown instruction", code)
			os.Exit(1)
		}
	}
}

func (vm *VM) binaryop(op func(int, int) int) {
	value1 := vm.pop()
	value2 := vm.pop()
	result := op(value1, value2)
	vm.push(result)
}

func (vm *VM) pop() int {
	vm.sp -= 1
	if vm.sp < 0 {
		fmt.Printf("Stack underflow: accessing sp at %d\n", vm.sp)
		os.Exit(1)
	}

	return vm.stack[vm.sp]
}

func (vm *VM) push(value int) {
	vm.stack[vm.sp] = value
	vm.sp += 1
}

func (vm *VM) LastPoppedInt() int {
	return vm.stack[vm.sp]
}

func boolToInt(value bool) int {
	if value {
		return 1
	}

	return 0
}
