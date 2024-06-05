package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hallo %s This is the monkey programming language\n", user.Username)

	fmt.Printf("Feel free to test commands\n")

	repl.Start(os.Stdin, os.Stdout)

}
