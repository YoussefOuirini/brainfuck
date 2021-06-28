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
}
