package cmd

import (
	"bufio"
	"fmt"
	"os"
)

const dataSize int = 65535

func ExecuteBf(program []Instruction) {
	data := make([]int16, dataSize)
	var dataPointer uint16 = 0
	reader := bufio.NewReader(os.Stdin)
	for pc := 0; pc < len(program); pc++ {
		switch program[pc].operator {
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
			read_val, _ := reader.ReadByte()
			data[dataPointer] = int16(read_val)
		case jumpForward:
			if data[dataPointer] == 0 {
				pc = int(program[pc].operand)
			}
		case jumpBackward:
			if data[dataPointer] > 0 {
				pc = int(program[pc].operand)
			}
		default:
			panic("Unknown operator.")
		}
	}
}
