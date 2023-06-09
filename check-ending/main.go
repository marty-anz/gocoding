package main

import (
	"fmt"
)

// https://edabit.com/challenge/PZnwXraqBPYv7w4Sm
func checkEnding(s string, e string) bool {
	sl := len(s)
	el := len(e)

	if el > sl {
		return false
	}

	for i := 1; i <= el; i++ {
		if e[el-i] != s[sl-i] {
			return false
		}
	}

	return true
}

func main() {
	inputs := [][2]string{
		{"abc", "bc"},
		{"abc", "d"},
		{"samurai", "zi"},
		{"feminine", "nine"},
		{"convention", "tio"},
	}

	for _, input := range inputs {
		s, e := input[0], input[1]
		fmt.Printf("checkEnding(%s, %s) -> %t\n", s, e, checkEnding(s, e))
	}
}
