package main

import (
	"fmt"
	"os"
	"os/user"

	"monkey/repl"
)

func main() {
	err := runRepl()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}

func runRepl() error {
	user, err := user.Current()
	if err != nil {
		return fmt.Errorf("couldn't determine user: %s", err.Error())
	}

	fmt.Printf("Hello %s! This is the Monkey programming language\n", user.Username)
	fmt.Print("Feel free to type in commands\n")

	return repl.Start(os.Stdin, os.Stdout)
}
