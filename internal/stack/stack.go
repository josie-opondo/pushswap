package stack

import (
	"errors"
)

// Stack represents a stack of integers.
type Stack struct {
	data []int
}

// NewStack initializes a new stack.
func NewStack(nums []int) *Stack {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return &Stack{nums}
}

// Push adds an element to the top of the stack.
func (s *Stack) PushFront(val int) {
	s.data = append([]int{val}, s.data...)
}

func (s *Stack) PushBack(n int) {
	s.data = append(s.data, n)
}

// Pop removes and returns the top element of the stack.
func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	res := s.data[0]
	s.data = s.data[1:]
	return res
}

func (s *Stack) PopBack() int {
	if len(s.data) == 0 {
		return -1
	}
	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return res
}

 //push the top first element of stack b to stack a
 func pa(s []*Stack) {
	if len(s[1].data) > 0 {
		s[0].PushBack(s[1].PopBack())
	}
}

 //push the top first element of stack a to stack b
 func pb(s []*Stack) {
	if len(s[0].data) > 0 {
		s[1].PushBack(s[0].PopBack())
	}
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

// ReverseRotate shifts all elements down by one position.
func (s *Stack) ReverseRotate() {
	if len(s.data) > 1 {
		last := s.data[len(s.data)-1]
		s.data = append([]int{last}, s.data[:len(s.data)-1]...)
	}
}

//swap first 2 elements of stack a
func sa(s []*Stack) {
	if len(s[0].data) > 1 {
		v1 := s[0].PopBack()
		v2 := s[0].PopBack()
		s[0].PushBack(v1)
		s[0].PushBack(v2)
	}
}

 //swap first 2 elements of stack b
func sb(s []*Stack) {
	if len(s[1].data) > 1 {
		v1 := s[1].PopBack()
		v2 := s[1].PopBack()
		s[1].PushBack(v1)
		s[1].PushBack(v2)
	}
}

 //execute sa and sb
func ss(s []*Stack) {
	sa(s)
	sb(s)
}

 //rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
func ra(s []*Stack) {
	if len(s[0].data) == 0 {
		return
	}
	s[0].PushFront(s[0].PopBack())
}

 //rotate stack b
func rb(s []*Stack) {
	if len(s[1].data) == 0 {
		return
	}
	s[1].PushFront(s[1].PopBack())
}

 //execute ra and rb
func rr(s []*Stack) {
	ra(s)
	rb(s)
}

//reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func rra(s []*Stack) {
	if len(s[0].data) == 0 {
		return
	}
	s[0].PushBack(s[0].Pop())
}

 //reverse rotate b
func rrb(s []*Stack) {
	if len(s[1].data) == 0 {
		return
	}
	s[1].PushBack(s[1].Pop())
}

 //execute rra and rrb
func rrr(s []*Stack) {
	rra(s)
	rrb(s)
}


type fn func(stack []*Stack)

var Cmd = map[string]fn{
	"pa":  pa,
	"pb":  pb,
	"sa":  sa,
	"sb":  sb,
	"ss":  ss,
	"ra":  ra,
	"rb":  rb,
	"rr":  rr,
	"rra": rra,
	"rrb": rrb,
	"rrr": rrr,
}

func Operation(stack []*Stack, instr string) {
	Cmd[instr](stack)
}

func RadixSort(data []int) string {
	res := ""
	stack := []*Stack{NewStack(data), NewStack(make([]int, 0, len(data)))}
	maxNum := len(stack[0].data) - 1
	maxBits := 0
	for maxNum>>maxBits != 0 {
		maxBits++
	}
	for i := 0; i < maxBits; i++ {
		for j := 0; j < len(data); j++ {
			num := stack[0].data[len(stack[0].data)-1]
			if (num>>i)&1 == 1 {
				Operation(stack, "ra")
				res += "ra\n"
			} else {
				Operation(stack, "pb")
				res += "pb\n"
			}
		}
		for len(stack[1].data) != 0 {
			Operation(stack, "pa")
			res += "pa\n"
		}
	}
	return res
}
