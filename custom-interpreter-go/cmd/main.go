package main

import (
	"custom-interpreter-go/pkg/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Banana programming language!\nFeel free to type in commands. Type `exit` command to exit.\n",
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
	fmt.Printf("Goodbye!\n")
}
