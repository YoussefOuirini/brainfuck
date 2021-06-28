package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/youssefouirini/brainfuck/cmd"
	"github.com/youssefouirini/brainfuck/model"
)

func Test_CompileBf(t *testing.T) {
	t.Run("runs with no error", func(t *testing.T) {
		_, err := cmd.CompileBf("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")
		assert.NoError(t, err)
	})

	t.Run("errors when no jumpstack", func(t *testing.T) {
		_, err := cmd.CompileBf("[")
		assert.EqualError(t, err, "compilation error: jumpStack is not 0")
	})
}

func Test_CompileProgram(t *testing.T) {
	t.Run("operation is nil", func(t *testing.T) {
		testChar := model.Char('k')
		testCounter := &model.Counter{Pointer: 1}
		err := cmd.CompileProgram(testChar, testCounter)
		assert.NoError(t, err)
		assert.Equal(t, uint16(1), testCounter.Pointer)
	})

	t.Run("operation is JumpForward", func(t *testing.T) {
		testChar := model.Char('[')
		testCounter := &model.Counter{Pointer: 1}
		err := cmd.CompileProgram(testChar, testCounter)
		assert.NoError(t, err)
		assert.Equal(t, []uint16{1}, testCounter.JumpStack)
	})

	t.Run("operation is JumpBackward errors", func(t *testing.T) {
		testChar := model.Char(']')
		testCounter := &model.Counter{JumpStack: []uint16{}}
		err := cmd.CompileProgram(testChar, testCounter)
		assert.EqualError(t, err, "compilation error: jumpStack is 0")
	})

	t.Run("operation is JumpBackward", func(t *testing.T) {
		testChar := model.Char(']')
		testCounter := &model.Counter{
			Pointer:     uint16(1),
			JumpPointer: uint16(0),
			JumpStack:   []uint16{1},
			Program:     []model.Instruction{{Operand: 1}, {Operand: 1}},
		}
		err := cmd.CompileProgram(testChar, testCounter)
		assert.NoError(t, err)
		assert.Equal(t, uint16(1), testCounter.JumpPointer)
		assert.Empty(t, testCounter.JumpStack)
		assert.Equal(t, uint16(1), testCounter.Program[0].Operand)
		assert.Equal(t, uint16(2), testCounter.Pointer)
	})

	t.Run("operation is IncreasePointer", func(t *testing.T) {
		testChar := model.Char('>')
		testCounter := &model.Counter{
			Pointer:     uint16(1),
			JumpPointer: uint16(0),
			JumpStack:   []uint16{1},
		}
		err := cmd.CompileProgram(testChar, testCounter)
		assert.NoError(t, err)
		assert.Equal(t, model.Instruction{Operator: model.IncreasePointer, Operand: 0}, testCounter.Program[0])
	})

	t.Run("operation is DecreasePointer", func(t *testing.T) {
		testChar := model.Char('<')
		testCounter := &model.Counter{
			Pointer:     uint16(1),
			JumpPointer: uint16(0),
			JumpStack:   []uint16{1},
		}
		err := cmd.CompileProgram(testChar, testCounter)
		assert.NoError(t, err)
		assert.Equal(t, model.Instruction{Operator: model.DecreasePointer, Operand: 0}, testCounter.Program[0])
	})

	t.Run("operation is IncreaseValue", func(t *testing.T) {
		testChar := model.Char('+')
		testCounter := &model.Counter{
			Pointer:     uint16(1),
			JumpPointer: uint16(0),
			JumpStack:   []uint16{1},
		}
		err := cmd.CompileProgram(testChar, testCounter)
		assert.NoError(t, err)
		assert.Equal(t, model.Instruction{Operator: model.IncreaseValue, Operand: 0}, testCounter.Program[0])
	})

	t.Run("operation is DecreaseValue", func(t *testing.T) {
		testChar := model.Char('-')
		testCounter := &model.Counter{
			Pointer:     uint16(1),
			JumpPointer: uint16(0),
			JumpStack:   []uint16{1},
		}
		err := cmd.CompileProgram(testChar, testCounter)
		assert.NoError(t, err)
		assert.Equal(t, model.Instruction{Operator: model.DecreaseValue, Operand: 0}, testCounter.Program[0])
	})

	t.Run("operation is Out", func(t *testing.T) {
		testChar := model.Char('.')
		testCounter := &model.Counter{
			Pointer:     uint16(1),
			JumpPointer: uint16(0),
			JumpStack:   []uint16{1},
		}
		err := cmd.CompileProgram(testChar, testCounter)
		assert.NoError(t, err)
		assert.Equal(t, model.Instruction{Operator: model.Out, Operand: 0}, testCounter.Program[0])
	})

	t.Run("operation is In", func(t *testing.T) {
		testChar := model.Char(',')
		testCounter := &model.Counter{
			Pointer:     uint16(1),
			JumpPointer: uint16(0),
			JumpStack:   []uint16{1},
		}
		err := cmd.CompileProgram(testChar, testCounter)
		assert.NoError(t, err)
		assert.Equal(t, model.Instruction{Operator: model.In, Operand: 0}, testCounter.Program[0])
	})
}
