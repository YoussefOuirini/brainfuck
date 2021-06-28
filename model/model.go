package model

import (
	"bufio"
	"os"
)

const dataSize int = 65535

const (
	Unknown = iota
	IncreasePointer
	DecreasePointer
	IncreaseValue
	DecreaseValue
	Out
	In
	JumpForward
	JumpBackward
)

type Counter struct {
	Pointer     uint16
	JumpPointer uint16
	JumpStack   []uint16
	Program     Program
}

type Char rune

func (c Char) GetOperation() *uint16 {
	var operation uint16

	switch c {
	case '>':
		operation = IncreasePointer
	case '<':
		operation = DecreasePointer
	case '+':
		operation = IncreaseValue
	case '-':
		operation = DecreaseValue
	case '.':
		operation = Out
	case ',':
		operation = In
	case '[':
		operation = JumpForward
	case ']':
		operation = JumpBackward
	}

	if operation == uint16(0) {
		return nil
	}

	return &operation
}

type Instruction struct {
	Operator uint16
	Operand  uint16
}

type Program []Instruction

func (p Program) Execute() string {
	data := make([]int16, dataSize)
	var dataPointer uint16 = 0

	var result []rune

	reader := bufio.NewReader(os.Stdin)
	for programCounter := 0; programCounter < len(p); programCounter++ {
		instruction := p[programCounter]
		switch instruction.Operator {
		case IncreasePointer:
			dataPointer++
		case DecreasePointer:
			dataPointer--
		case IncreaseValue:
			data[dataPointer]++
		case DecreaseValue:
			data[dataPointer]--
		case Out:
			result = append(result, rune(data[dataPointer]))
		case In:
			read_val, err := reader.ReadByte()
			if err != nil {
				panic(err)
			}
			data[dataPointer] = int16(read_val)
		case JumpForward:
			if data[dataPointer] == 0 {
				programCounter = int(instruction.Operand)
			}
		case JumpBackward:
			if data[dataPointer] > 0 {
				programCounter = int(instruction.Operand)
			}
		default:
			panic("unknown operator")
		}
	}

	return string(result)
}
