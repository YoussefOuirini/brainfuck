package cmd

import (
	"bufio"
	"fmt"
	"os"
)

const dataSize int = 65535

type Program []Instruction

func (p Program) ExecuteProgram() {
	data := make([]int16, dataSize)
	var dataPointer uint16 = 0

	reader := bufio.NewReader(os.Stdin)
	for pc := 0; pc < len(p); pc++ {
		switch p[pc].operator {
		case increasePointer:
			dataPointer++
		case decreasePointer:
			dataPointer--
		case increaseValue:
			data[dataPointer]++
		case decreaseValue:
			data[dataPointer]--
		case out:
			fmt.Printf("%c", data[dataPointer])
		case in:
			read_val, err := reader.ReadByte()
			if err != nil {
				panic(err)
			}
			data[dataPointer] = int16(read_val)
		case jumpForward:
			if data[dataPointer] == 0 {
				pc = int(p[pc].operand)
			}
		case jumpBackward:
			if data[dataPointer] > 0 {
				pc = int(p[pc].operand)
			}
		default:
			panic("Unknown operator.")
		}
	}
}
