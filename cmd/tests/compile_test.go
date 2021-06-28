package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/youssefouirini/brainfuck/cmd"
)

func Test_CompileBf(t *testing.T) {
	t.Run("runs with no error", func(t *testing.T) {
		program, err := cmd.CompileBf("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")
		t.Log(program)
		assert.NoError(t, err)
	})
}
