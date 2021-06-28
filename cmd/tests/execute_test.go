package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/youssefouirini/brainfuck/cmd"
	"github.com/youssefouirini/brainfuck/model"
)

func Test_ExecuteBf(t *testing.T) {
	t.Run("succes runs Hello World!", func(t *testing.T) {
		testContent := model.Brainfuck{Contents: []byte("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")}
		result, err := cmd.ExecuteBf(testContent)
		assert.NoError(t, err)
		assert.Equal(t, "Hello World!\n", result)
	})

	t.Run("666", func(t *testing.T) {
		testContent := model.Brainfuck{Contents: []byte(">+++++++++[<++++++>-]<...>++++++++++.")}
		result, err := cmd.ExecuteBf(testContent)
		assert.NoError(t, err)
		assert.Equal(t, "666\n", result)
	})
}
