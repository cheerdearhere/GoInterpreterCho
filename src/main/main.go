package main

import (
	"GoInterpreter/src/main/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commaand\n")
	repl.Start(os.Stdin, os.Stdout)
}
