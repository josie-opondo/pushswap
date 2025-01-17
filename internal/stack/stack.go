package stack

import (
	"errors"
	"fmt"
)

// Stack represents a stack of integers.
type Stack struct {
	data []int
	name string
}

// NewStack initializes a new stack.
func NewStack(name string) *Stack {
	return &Stack{data: []int{}, name: name}
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(val int) {
	s.data = append([]int{val}, s.data...)
}