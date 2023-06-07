package main

import (
	"fmt"
)

// https://edabit.com/challenge/5LnycSd2xT4uwZCpi

func countWords(s string) uint {
	total := 0
	inWord := false

	for _, c := range s {
		if c != ' ' {
			inWord = true
		} else if inWord {
			total += 1
			inWord = false
		}
	}

	if inWord {
		total += 1
	}

	return uint(total)
}

func main() {
	inputs := []string{
		"Just an example here move along",
		"This is a test",
		"What an easy task, right",
		"What an   easy task, right  ",
	}

	for _, input := range inputs {
		fmt.Printf("%s -> %d\n", input, countWords(input))
	}
}
