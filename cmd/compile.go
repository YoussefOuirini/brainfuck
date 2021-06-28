package cmd

import (
	"errors"
)

type Instruction struct {
	operator uint16
	operand  uint16
}

type Counter struct {
	program uint16
	jump    uint16
	stack   []uint16
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

func CompileBf(input string) (program []Instruction, err error) {
	var counter Counter
	counter.stack = make([]uint16, 0)
	for _, char := range input {
		instruction, err := compileProgram(&counter, char)
		if err != nil {
			return nil, err
		}

		if instruction != nil {
			program = append(program, *instruction)
		}

		if instruction.operator == jumpBackward {
			program[counter.jump].operand = counter.jump
		}
	}
	if len(counter.stack) != 0 {
		return nil, errors.New("Compilation error.")
	}
	return
}

func compileProgram(counter *Counter, char rune) (*Instruction, error) {
	instruction := compileChar(char)
	counter.program++

	if instruction == nil {
		counter.jump--
	}

	if instruction.operator == jumpForward {
		counter.stack = append(counter.stack, counter.program)
	}

	if instruction.operator == jumpBackward {
		if len(counter.stack) == 0 {
			return nil, errors.New("compilation error due to jumpStack being 0")
		}

		counter.jump = counter.stack[len(counter.stack)-1]
		counter.stack = counter.stack[:len(counter.stack)-1]
		instruction.operand = counter.jump
	}

	return instruction, nil
}

func compileChar(char rune) *Instruction {
	var operator uint16
	switch char {
	case '>':
		operator = increasePointer
	case '<':
		operator = decreasePointer
	case '+':
		operator = increaseValue
	case '-':
		operator = decreaseValue
	case '.':
		operator = out
	case ',':
		operator = in
	case '[':
		operator = jumpForward
	case ']':
		operator = jumpBackward
	default:
		return nil
	}

	return &Instruction{operator, 0}
}
