package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")                                // prompt
		command, err := bufio.NewReader(os.Stdin).ReadString('\n') // Wait for user input
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		fmt.Println(command[:len(command)-1] + ": command not found") // strip newline
	}
}
