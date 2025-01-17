package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	stk "pushswap/internal/stack"
	"sort"
	"strings"
	"sync"
)

type Checker struct{}

var (
	instance *Checker
	once     sync.Once
)

func GetCheckerInstance() *Checker {
	once.Do(func() {
		instance = &Checker{}
	})
	return instance
}

func (c *Checker) GetInput() ([]byte, error) {
	if stat, err := os.Stdin.Stat(); err != nil || (stat.Mode()&os.ModeNamedPipe == 0) {
		return nil, fmt.Errorf("Error")
	}

	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("Error")
	}
	return content, nil
}

func (c *Checker) GetOutput(instructions []byte, stack []*stk.Stack) (string, error) {
	operations := strings.Split(string(instructions), "\n")
	l := len(operations) - 1
	if l > 0 && operations[l-1] == "" {
		l--
	}
	operations = operations[:l]
	for _, v := range operations {
		if f, ok := stk.Cmd[v]; ok {
			f(stack)
		} else {
			return "", fmt.Errorf("Error")
		}
	}

	res := stack[0].Data
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	if sort.IntsAreSorted(res) {
		return "OK", nil
	} else {
		return "KO", nil
	}
}
