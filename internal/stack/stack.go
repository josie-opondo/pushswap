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

