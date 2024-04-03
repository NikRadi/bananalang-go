package vm

import (
	"fmt"
	"os"
)

type (
	LocalVariables []int

	OperandStack []int

	Frame struct {
		LocalVariables LocalVariables
		OperandStack   OperandStack
	}

	CallStack []*Frame
)

func (stack *CallStack) current() *Frame {
	index := len(*stack) - 1
	if index < 0 {
		fmt.Printf("Stack underflow: accessing stack at %d\n", index)
		os.Exit(1)
	}

	frame := (*stack)[index]
	return frame
}

func (stack *CallStack) pop() *Frame {
	index := len(*stack) - 1
	frame := (*stack)[index]
	*stack = (*stack)[:index]
	return frame
}

func (stack *CallStack) push(frame *Frame) {
	*stack = append(*stack, frame)
}

func (frame *Frame) pop() int {
	index := len(frame.OperandStack) - 1
	value := frame.OperandStack[index]
	frame.OperandStack = frame.OperandStack[:index]
	return value
}

func (frame *Frame) push(value int) {
	frame.OperandStack = append(frame.OperandStack, value)
}

func (frame *Frame) load(index int) int {
	return frame.LocalVariables[index]
}

func (frame *Frame) store(index, value int) {
	frame.LocalVariables[index] = value
}
