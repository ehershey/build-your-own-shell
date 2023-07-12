package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		// Read the keyboad input.
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				os.Exit(0)
			}
			// fmt.Fprintln(os.Stderr, err)
			fmt.Fprintf(os.Stderr, "got error reading input! (%v)\n", err)
		}
		fmt.Fprintf(os.Stderr, "input was: %v", input)
		err = execInput(input)
		if err != nil {
			// fmt.Fprintln(os.Stderr, err)
			fmt.Fprintf(os.Stderr, "got error executing command! (%v)\n", err)
		}
	}
}

func execInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		var dirname string
		var err error
		if len(args) == 1 || (len(args) == 2 && args[1] == "~") {
			dirname, err = os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("error getting home directry: %v", err)
			}
		} else {
			dirname = args[1]
		}
		err = os.Chdir(dirname)
		if err != nil {
			return fmt.Errorf("error changing directory: %v", err)
		}
	case "exit":
		os.Exit(0)
	}

	// Pass the program and the arguments separately.
	cmd := exec.Command(args[0], args[1:]...)

	// Prepare the command to execute.

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}
