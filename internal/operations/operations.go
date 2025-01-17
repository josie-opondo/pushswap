package operations

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func FlagRandoms() bool {
	var random = flag.Bool("rand", false, "Use -rand to be able to generate random numbers")
	flag.Parse()
	return *random
}

func InputNums() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return GetInts(scanner.Text())
}

func GenerateRandomNums() ([]int, error) {
	fmt.Println("Numbers will be randomly generated in the range [min,max).\nEnter the minimum, maximum values and the length of numbers in order.\nmin max length")
	input := InputNums()
	if len(input) != 3 {
		return nil, fmt.Errorf("please, check your input for correctness")
	}
	return UniqRandInts(input[0], input[1], input[2])
}

func GetInts(s string) []int {
	fields := strings.Fields(s)
	nums := make([]int, 0, len(fields))
	for _, f := range fields {
		n, err := strconv.Atoi(f)
		if err == nil {
			nums = append(nums, n)
		}
	}
	return nums
}

func UniqInts(s string) ([]int, error) {
	input := strings.Split(s, " ")
	nums := make([]int, len(input))
	existing := make(map[int]bool, len(nums))
	for i, n := range input {
		num, err := strconv.Atoi(n)
		if err != nil || existing[num] {
			return nil, fmt.Errorf("invalid number")
		}
		existing[num] = true
		nums[i] = num
	}
	return nums, nil
}

func UniqRandInts(min, max, length int) ([]int, error) {
	if min > max {
		return nil, fmt.Errorf("minimum number can't be larger than %v", max)
	}
	if max-min < length {
		return nil, fmt.Errorf("numbers can't be duplicated, try length >= max-min")
	}
	nums := make([]int, length)
	randomList := rand.Perm(max - min + 1)
	for i := 0; i < length; i++ {
		nums[i] = randomList[i] + min
	}
	return nums, nil
}
