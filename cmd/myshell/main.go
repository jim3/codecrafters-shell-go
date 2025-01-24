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
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		// Split up `command` so that you know what you are executing
		input := strings.Fields(strings.TrimSpace(command))

		// Terminate wtih exit status 0 if the command is `exit 0`
		if input[0] == "exit" && len(input) == 2 && input[1] == "0" {
			os.Exit(0)
		}

		// Execute the command
		fmt.Printf("%s: command not found\n", strings.TrimSpace(command))
	}
}
