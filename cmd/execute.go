package cmd

import "github.com/youssefouirini/brainfuck/model"

func ExecuteBf(bf model.Brainfuck) (string, error) {
	program, err := CompileBf(string(bf.Contents))
	if err != nil {
		return "", err
	}
	return program.Execute(), nil
}
