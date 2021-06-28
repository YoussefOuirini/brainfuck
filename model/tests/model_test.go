package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/youssefouirini/brainfuck/model"
)

func Test_GetOperation(t *testing.T) {
	t.Run("IncreasePointer", func(t *testing.T) {
		testChar := model.Char('>')
		operation := testChar.GetOperation()
		assert.Equal(t, uint16(model.IncreasePointer), operation)
	})

	t.Run("DecreasePointer", func(t *testing.T) {
		testChar := model.Char('<')
		operation := testChar.GetOperation()
		assert.Equal(t, uint16(model.DecreasePointer), operation)
	})

	t.Run("IncreaseValue", func(t *testing.T) {
		testChar := model.Char('+')
		operation := testChar.GetOperation()
		assert.Equal(t, uint16(model.IncreaseValue), operation)
	})

	t.Run("DecreaseValue", func(t *testing.T) {
		testChar := model.Char('-')
		operation := testChar.GetOperation()
		assert.Equal(t, uint16(model.DecreaseValue), operation)
	})

	t.Run("In", func(t *testing.T) {
		testChar := model.Char(',')
		operation := testChar.GetOperation()
		assert.Equal(t, uint16(model.In), operation)
	})

	t.Run("Out", func(t *testing.T) {
		testChar := model.Char('.')
		operation := testChar.GetOperation()
		assert.Equal(t, uint16(model.Out), operation)
	})

	t.Run("JumpForward", func(t *testing.T) {
		testChar := model.Char('[')
		operation := testChar.GetOperation()
		assert.Equal(t, uint16(model.JumpForward), operation)
	})

	t.Run("JumpBackward", func(t *testing.T) {
		testChar := model.Char(']')
		operation := testChar.GetOperation()
		assert.Equal(t, uint16(model.JumpBackward), operation)
	})

	t.Run("Unknown", func(t *testing.T) {
		testChar := model.Char('t')
		operation := testChar.GetOperation()
		assert.Equal(t, uint16(model.Unknown), operation)
	})
}

func Test_Execute(t *testing.T) {
	t.Run("panics when unknown operator", func(t *testing.T) {
		testProgram := model.Program{
			{
				Operator: 100,
			},
		}

		panickingProgram := func() {
			testProgram.Execute()
		}

		assert.PanicsWithValue(t, "unknown operator", panickingProgram)
	})
}
