package cmd

import (
	"errors"

	"github.com/youssefouirini/brainfuck/model"
)

func CompileBf(input string) (model.Program, error) {
	counter := &model.Counter{}
	counter.JumpStack = make([]uint16, 0)
	for _, char := range input {
		CompileProgram(model.Char(char), counter)
	}
	if len(counter.JumpStack) != 0 {
		return nil, errors.New("compilation error: jumpStack is not 0")
	}
	return counter.Program, nil
}

func CompileProgram(char model.Char, counter *model.Counter) error {
	operation := char.GetOperation()

	var operand uint16

	if operation == model.Unknown {
		counter.Pointer--
	}

	if operation == model.JumpForward {
		counter.JumpStack = append(counter.JumpStack, counter.Pointer)
	}

	if operation == model.JumpBackward {
		if len(counter.JumpStack) == 0 {
			return errors.New("compilation error: jumpStack is 0")
		}
		counter.JumpPointer = counter.JumpStack[len(counter.JumpStack)-1]
		counter.JumpStack = counter.JumpStack[:len(counter.JumpStack)-1]
		counter.Program[counter.JumpPointer].Operand = counter.Pointer
		operand = counter.JumpPointer
	}

	counter.Program = append(counter.Program, model.Instruction{Operator: operation, Operand: operand})

	counter.Pointer++

	return nil
}
