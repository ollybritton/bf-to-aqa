package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("There was an issue getting information about stdin:")
		fmt.Println("Error: ", err)
		return
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Print("This program works with pipes. Please pipe your data like so:\n\n")
		fmt.Println("  cat program.bf | bf-to-aqa")
		return
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("There was an issue getting text from stdin:")
		fmt.Println("Error: ", err)
		return
	}

	instructions := createInstructions(bytes)
	program := assemble(instructions)
	simple := simplify(program)
	fmt.Println(simple)
}
