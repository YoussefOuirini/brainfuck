package cmd

import "errors"

type Instruction struct {
	operator uint16
	operand  uint16
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
	var programCounter, jumpProgramCounter uint16 = 0, 0
	jumpStack := make([]uint16, 0)
	for _, char := range input {
		instruction := compileChar(char, programCounter, jumpProgramCounter, jumpStack, program)
		programCounter++

		if instruction == nil {
			programCounter--
		}

		if instruction.operator == jumpForward {
			jumpStack = append(jumpStack, programCounter)
		}

		if instruction.operator == jumpBackward {
			if len(jumpStack) == 0 {
				return nil, errors.New("compilation error due to jumpStack being 0")
			}

			jumpProgramCounter = jumpStack[len(jumpStack)-1]
			jumpStack = jumpStack[:len(jumpStack)-1]
			program = append(program, Instruction{jumpBackward, jumpProgramCounter})
			program[jumpProgramCounter].operand = programCounter

			continue
		}

		program = append(program, *instruction)
	}
	if len(jumpStack) != 0 {
		return nil, errors.New("Compilation error.")
	}
	return
}

func compileChar(char rune, programCounter, jumpProgramCounter uint16, jumpStack []uint16, program []Instruction) *Instruction {
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
