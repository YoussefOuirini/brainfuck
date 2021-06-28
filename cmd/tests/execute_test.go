package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/youssefouirini/brainfuck/cmd"
)

func Test_ExecuteBf(t *testing.T) {
	t.Run("succes runs Hello World!", func(t *testing.T) {
		testContent := []byte("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")
		result, err := cmd.ExecuteBf(testContent)
		assert.NoError(t, err)
		assert.Equal(t, "Hello World!\n", result)
	})
}
