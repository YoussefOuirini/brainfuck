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

	result, err := cmd.ExecuteBf([]byte(args[1]))
	if err != nil {
		fmt.Printf("Error executing program %s\n", err)
	}

	fmt.Print(result)
}
