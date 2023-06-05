package main

import (
	"fmt"
)

// source: https://edabit.com/challenge/6MZx5RqKYkFaogeAQ

func getOnlyEvens(lst []int) (evens []int) {
	for i := 0; i < len(lst); i += 2 {
		if lst[i]%2 == 0 {
			evens = append(evens, lst[i])
		}
	}

	return
}

func main() {
	inputs := [][]int{
		{1, 3, 2, 6, 4, 8},
		{0, 1, 2, 3, 4},
		{1, 2, 3, 4, 5},
	}

	for _, input := range inputs {
		fmt.Printf("%v -> %v\n", input, getOnlyEvens(input))
	}
}
