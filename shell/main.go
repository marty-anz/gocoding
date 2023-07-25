package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// a simple shell demostration program
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			panic(err)
		}

		line := scanner.Text() // a command line
		cmd := strings.Split(line, " ")

		switch cmd[0] { // the first word is the cmmand, and the rest are arguments
		case "echo": // built-in command
			{
				fmt.Println(strings.Join(cmd[1:], " "))
			}
		case "bye": // built-in command
			{
				fmt.Println("bye bye")
				return
			}
		default: // otherwise, try to execute the program
			{
				prog := exec.Command(cmd[0], cmd[1:]...)
				out, err := prog.Output()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(string(out))
				}
			}
		}
	}
}
