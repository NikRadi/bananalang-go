package vm

import (
	"bananalang/opcode"
	"fmt"
	"os"
)

type VM struct {
	Stack CallStack
}

func NewVM() *VM {
	frame := &Frame{
		LocalVariables: make(LocalVariables, 0),
		OperandStack:   make(OperandStack, 0),
	}

	stack := make(CallStack, 0)
	stack.push(frame)
	return &VM{Stack: stack}
}

func (vm *VM) Execute(codes []opcode.Opcode) {
	for i := 0; i < len(codes); i += 1 {
		code := codes[i]
		switch code {
		case opcode.Push:
			i += 1
			value := int(codes[i])
			vm.Stack.current().push(value)
		case opcode.Pop:
			vm.Stack.current().pop()
		case opcode.Print:
			value := vm.Stack.current().pop()
			fmt.Println(value)
		case opcode.Add:
			vm.binaryop(func(v1, v2 int) int { return v2 + v1 })
		case opcode.Sub:
			vm.binaryop(func(v1, v2 int) int { return v2 - v1 })
		case opcode.Mul:
			vm.binaryop(func(v1, v2 int) int { return v2 * v1 })
		case opcode.Neg:
			value := vm.Stack.current().pop()
			vm.Stack.current().push(-value)
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
	value1 := vm.Stack.current().pop()
	value2 := vm.Stack.current().pop()
	result := op(value1, value2)
	vm.Stack.current().push(result)
}

func boolToInt(value bool) int {
	if value {
		return 1
	}

	return 0
}
