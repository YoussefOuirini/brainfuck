package cmd

import (
	"bufio"
	"os"
)

const dataSize int = 65535

type Program []Instruction

type DataPointer uint16

func (p Program) ExecuteProgram() string {
	data := make([]int16, dataSize)
	var dataPointer DataPointer = 0

	var result []rune

	reader := bufio.NewReader(os.Stdin)
	for programCounter := 0; programCounter < len(p); programCounter++ {
		switch p[programCounter].operator {
		case increasePointer:
			dataPointer++
		case decreasePointer:
			dataPointer--
		case increaseValue:
			data[dataPointer]++
		case decreaseValue:
			data[dataPointer]--
		case out:
			result = append(result, rune(data[dataPointer]))
		case in:
			read_val, err := reader.ReadByte()
			if err != nil {
				panic(err)
			}
			data[dataPointer] = int16(read_val)
		case jumpForward:
			if data[dataPointer] == 0 {
				programCounter = int(p[programCounter].operand)
			}
		case jumpBackward:
			if data[dataPointer] > 0 {
				programCounter = int(p[programCounter].operand)
			}
		default:
			panic("Unknown operator.")
		}
	}

	return string(result)
}
