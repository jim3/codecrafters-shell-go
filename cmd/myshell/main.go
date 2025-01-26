package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	PROMPT = "$ "
)

func main() {
	for {
		fmt.Fprint(os.Stdout, PROMPT)

		command, err := bufio.NewReader(os.Stdin).ReadString('\n') // -> (string, error)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		input := strings.Fields(strings.TrimSpace(command))
		// Check if there's any input
		if len(input) > 0 {
			if input[0] == "exit" && input[1] == "0" {
				os.Exit(0)
			}

			// implement echo cmd
			if input[0] == "echo" {
				echoSlice := input[1:]
				strEcho := strings.Join(echoSlice, " ")
				echo := strings.TrimSpace(strEcho)
				fmt.Println(echo)
			} else {
				fmt.Println(input[0] + ": command not found")
			}
		}
	}
}
