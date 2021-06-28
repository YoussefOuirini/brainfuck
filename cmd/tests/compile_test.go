package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/youssefouirini/brainfuck/cmd"
)

func Test_CompileBf(t *testing.T) {
	t.Run("runs with no error", func(t *testing.T) {
		_, err := cmd.CompileBf("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")
		assert.NoError(t, err)
	})

	t.Run("errors when no jumpstack", func(t *testing.T) {
		_, err := cmd.CompileBf("[")
		assert.Error(t, err)
		assert.EqualError(t, err, "compilation error: jumpStack is not 0")
	})
}
