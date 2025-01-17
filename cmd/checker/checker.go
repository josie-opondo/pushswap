package main

import (
	"fmt"
	"os"
	"pushswap/internal/helpers"
	operations "pushswap/internal/operations"
	stk "pushswap/internal/stack"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}

	c := helpers.GetCheckerInstance()

	nums, err := operations.UniqInts(args[0])
	if err != nil {
		fmt.Println("Error")
		return
	}

	stack := []*stk.Stack{stk.NewStack(nums), stk.NewStack(make([]int, 0, len(nums)))}

	instructions, err := c.GetInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	output, err := c.GetOutput(instructions, stack)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(output)
}
