package cmd

import (
	"errors"
)

type Instruction struct {
	operator uint16
	operand  uint16
}

type Counter struct {
	pointer     uint16
	jumpPointer uint16
	jumpStack   []uint16
	program     []Instruction
}

const (
	increasePointer = iota
	decreasePointer
	increaseValue
	decreaseValue
	out
	in
	jumpForward
	jumpBackward
)

func CompileBf(input string) ([]Instruction, error) {
	counter := &Counter{}
	counter.jumpStack = make([]uint16, 0)
	for _, char := range input {
		compileProgram(counter, char)
	}
	if len(counter.jumpStack) != 0 {
		return nil, errors.New("Compilation error.")
	}
	return counter.program, nil
}

func compileProgram(counter *Counter, char rune) error {
	switch char {
	case '>':
		counter.program = append(counter.program, Instruction{increasePointer, 0})
	case '<':
		counter.program = append(counter.program, Instruction{decreasePointer, 0})
	case '+':
		counter.program = append(counter.program, Instruction{increaseValue, 0})
	case '-':
		counter.program = append(counter.program, Instruction{decreaseValue, 0})
	case '.':
		counter.program = append(counter.program, Instruction{out, 0})
	case ',':
		counter.program = append(counter.program, Instruction{in, 0})
	case '[':
		counter.program = append(counter.program, Instruction{jumpForward, 0})
		counter.jumpStack = append(counter.jumpStack, counter.pointer)
	case ']':
		if len(counter.jumpStack) == 0 {
			return errors.New("Compilation error.")
		}
		counter.jumpPointer = counter.jumpStack[len(counter.jumpStack)-1]
		counter.jumpStack = counter.jumpStack[:len(counter.jumpStack)-1]
		counter.program = append(counter.program, Instruction{jumpBackward, counter.jumpPointer})
		counter.program[counter.jumpPointer].operand = counter.pointer
	default:
		counter.pointer--
	}
	counter.pointer++

	return nil
}
