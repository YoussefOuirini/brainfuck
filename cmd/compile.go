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

type Char rune

func (c Char) GetOperation() *uint16 {
	var operation uint16

	switch c {
	case '>':
		operation = increasePointer
	case '<':
		operation = decreasePointer
	case '+':
		operation = increaseValue
	case '-':
		operation = decreaseValue
	case '.':
		operation = out
	case ',':
		operation = in
	case '[':
		operation = jumpForward
	case ']':
		operation = jumpBackward
	}

	return &operation
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
		compileProgram(Char(char), counter)
	}
	if len(counter.jumpStack) != 0 {
		return nil, errors.New("compilation error: jumpStack is not 0")
	}
	return counter.program, nil
}

func compileProgram(char Char, counter *Counter) error {
	operation := char.GetOperation()

	var operand uint16

	if operation == nil {
		counter.pointer--
	}

	if *operation == jumpForward {
		counter.jumpStack = append(counter.jumpStack, counter.pointer)
	}

	if *operation == jumpBackward {
		if len(counter.jumpStack) == 0 {
			return errors.New("compilation error: jumpStack is 0")
		}
		counter.jumpPointer = counter.jumpStack[len(counter.jumpStack)-1]
		counter.jumpStack = counter.jumpStack[:len(counter.jumpStack)-1]
		counter.program[counter.jumpPointer].operand = counter.pointer
		operand = counter.jumpPointer
	}

	counter.program = append(counter.program, Instruction{*operation, operand})

	counter.pointer++

	return nil
}
