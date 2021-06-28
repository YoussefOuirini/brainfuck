package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/youssefouirini/brainfuck/cmd"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s filename\n", args[0])
		return
	}
	filename := args[1]
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %s\n", filename)
		return
	}

	result, err := cmd.ExecuteBf(fileContents)
	if err != nil {
		fmt.Printf("Error executing program %s\n", err)
	}

	fmt.Print(result)
}
