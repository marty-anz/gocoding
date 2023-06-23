package main

import (
	"fmt"
)

type step struct {
	from string
	to   string
}

func hanoi(discs uint, from, via, to string) (steps []*step) {
	if discs == 0 {
		return steps
	}

	subs := hanoi(discs-1, from, to, via)

	steps = append(steps, subs...)
	steps = append(steps, &step{from: from, to: to})

	subs = hanoi(discs-1, via, from, to)

	return append(steps, subs...)
}

func displaySteps(steps []*step) {
	for _, s := range steps {
		fmt.Printf("%s -> %s\n", s.from, s.to)
	}
}

func main() {
	displaySteps(hanoi(3, "a", "b", "c"))
}
