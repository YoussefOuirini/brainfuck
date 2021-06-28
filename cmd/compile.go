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
		err = compileProgram(programCounter, jumpProgramCounter, char, jumpStack, program)
		if err != nil {
			return nil, err
		}
	}
	if len(jumpStack) != 0 {
		return nil, errors.New("Compilation error.")
	}
	return
}

func compileProgram(programCounter, jumpProgramCounter uint16, char rune, jumpStack []uint16, program []Instruction) error {
	instruction := compileChar(char)
	programCounter++

	if instruction == nil {
		programCounter--
	}

	if instruction.operator == jumpForward {
		jumpStack = append(jumpStack, programCounter)
	}

	if instruction.operator == jumpBackward {
		if len(jumpStack) == 0 {
			return errors.New("compilation error due to jumpStack being 0")
		}

		jumpProgramCounter = jumpStack[len(jumpStack)-1]
		jumpStack = jumpStack[:len(jumpStack)-1]
		instruction.operand = jumpProgramCounter
		program[jumpProgramCounter].operand = programCounter
	}

	program = append(program, *instruction)

	return nil
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
