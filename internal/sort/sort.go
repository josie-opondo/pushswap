package sort

import (
	"sort"
	stk "pushswap/internal/stack"
)

func SmallSort(nums []int) string {
	res := ""
	stack := []*stk.Stack{stk.NewStack(nums), stk.NewStack(make([]int, 0, len(nums)))}
	res = SmallInstructions(nums, stack)
	return res
}

func SmallInstructions(nums []int, s []*stk.Stack) string {
	half := len(nums) / 2
	res := ""
	for {
		a, b := s[0].Data, s[1].Data
		if isASorted(a) {
			break
		}
		op := match(aStack(a, half), a, b)
		stk.Operation(s, op)
		res += op + "\n"
	}
	for {
		a, b := s[0].Data, s[1].Data
		if len(a) == len(nums) && isASorted(a) {
			break
		}
		op := bStack(b)
		stk.Operation(s, op)
		res += op + "\n"
	}
	return res
}

// sa, ra, rra, pb
func aStack(a []int, half int) string {
	swap := saIdx(a)
	mid, i := len(a)/2, len(a)-1
	if a[i] < half {
		return "pb"
	}
	if swap > 0 {
		if inOrder(a[i-1], a[i]) {
			return "sa"
		} else if swap < mid {
			return "rra"
		}
	} else {
		minNumIdx := smallest(a)
		if minNumIdx == i {
			return "pb"
		} else if minNumIdx < mid || inOrder(a[i], a[0]) {
			return "rra"
		}
	}
	return "ra"
}

// sb, rb, rrb, pa
func bStack(bstack []int) string {
	b := make([]int, len(bstack))
	copy(b, bstack)
	UInts(b)
	swap := sbIdx(b)
	mid, j := len(b)/2, len(b)-1
	if swap > 0 {
		if inOrder(b[j], b[j-1]) {
			return "sb"
		} else if swap < mid {
			return "rrb"
		}
	} else {
		maxNumIdx := biggest(b)
		if maxNumIdx == j {
			return "pa"
		} else if maxNumIdx < mid || inOrder(b[0], b[j]) {
			return "rrb"
		}
	}
	return "rb"
}

func match(ins string, a, b []int) string {
	if len(b) < 2 || isBSorted(b) {
		return ins
	}
	bins := bStack(b)
	switch ins { //ins is an instruction from aStack(a)
	case "pb":
		if bins == "sb" {
			if t := len(a) - saIdx(a); t < 3 {
				if t == 1 {
					return "ss"
				}
				return "ra"
			} else {
				return "sb"
			}
		}
	case "sa":
		if t := len(b) - sbIdx(b); t < 3 {
			if bins == "sb" {
				return "ss"
			}
			return "rb"
		}
	case "ra":
		if bins == "rb" || bins == "sb" && len(b) == 2 {
			return "rr"
		}
	case "rra":
		if bins == "rrb" || bins == "sb" && len(b) == 2 {
			return "rrr"
		}
	}
	return ins
}

func smallest(nums []int) int {
	min, idx := 0x3f3f3f3f, -1
	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
			idx = i
		}
	}
	return idx
}

func biggest(nums []int) int {
	max, idx := -1, -1
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
			idx = i
		}
	}
	return idx
}

func saIdx(nums []int) int {
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i]-nums[i-1] == 1 {
			return i
		}
	}
	return -1
}

func sbIdx(nums []int) int {
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i-1]-nums[i] == 1 {
			return i
		}
	}
	return -1
}

// true if a+1 = b
func inOrder(a, b int) bool {
	return b-a == 1
}

func isASorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]-nums[i+1] != 1 {
			return false
		}
	}
	return true
}

func isBSorted(nums []int) bool {
	return sort.IntsAreSorted(nums)
}

func UInts(nums []int) error {
	copiedNums := make([]int, len(nums))
	copy(copiedNums, nums)
	sort.Ints(copiedNums)

	for i, n := range nums {
		nums[i] = BinarySearch(copiedNums, n)
	}

	return nil
}

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
