package main

import (
	"fmt"
	"os"
	operations "pushswap/internal/operations"
	srt "pushswap/internal/sort"
	stk "pushswap/internal/stack"
	"sort"
)

var (
	nums         []int
	err          error
	instructions string
)

func main() {
	if operations.FlagRandoms() {
		nums, err = operations.GenerateRandomNums()
		if err != nil {
			fmt.Println("Error")
			return
		}
	} else {
		args := os.Args[1:]
		if len(args) < 1 {
			return
		}
		nums, err = operations.UniqInts(args[0])
		if err != nil {
			fmt.Println("Error")
			return
		}
	}

	if sort.IntsAreSorted(nums) {
		return
	}

	filepath, err := os.Create("push-swap-result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filepath.Close()

	str := fmt.Sprint(nums)
	_, err = filepath.WriteString("\"" + str[1:len(str)-1] + "\"\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = srt.UInts(nums)
	if err != nil {
		fmt.Println("Error")
		return
	}

	if len(nums) <= 100 {
		instructions = srt.SmallSort(nums)
	} else {
		instructions = stk.RadixSort(nums)
	}

	fmt.Print(instructions)
	_, err = filepath.WriteString(instructions)
	if err != nil {
		fmt.Println(err)
		return
	}
}
