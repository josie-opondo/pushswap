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

// Pop removes and returns the top element of the stack.
func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("stack is empty")
	}
	val := s.data[0]
	s.data = s.data[1:]
	return val, nil
}

// Swap swaps the first two elements of the stack.
func (s *Stack) Swap() error {
	if len(s.data) < 2 {
		return errors.New("not enough elements to swap")
	}
	s.data[0], s.data[1] = s.data[1], s.data[0]
	return nil
}

// Rotate shifts all elements up by one position.
func (s *Stack) Rotate() {
	if len(s.data) > 1 {
		first := s.data[0]
		s.data = append(s.data[1:], first)
	}
}