package main

import (
	"fmt"
	"os"

	"github.com/youssefouirini/brainfuck/cmd"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s filename\n", args[0])
		return
	}
	// filename := args[1]
	// fileContents, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	fmt.Printf("Error reading %s\n", filename)
	// 	return
	// }

	program, err := cmd.CompileBf(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd.ExecuteBf(program)
}
