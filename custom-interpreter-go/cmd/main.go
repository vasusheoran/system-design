package main

import (
	"custom-interpreter-go/pkg/repl"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello! This is the Banana programming language!\nFeel free to type in commands. Type `exit` command to exit.\n")
	repl.Start(os.Stdin, os.Stdout)
	fmt.Printf("Goodbye!\n")
}
