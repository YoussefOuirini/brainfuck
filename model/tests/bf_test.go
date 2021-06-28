package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/youssefouirini/brainfuck/model"
)

func Test_Add(t *testing.T) {
	bf := model.Brainfuck{Contents: []byte{'+'}}

	t.Run("test Add", func(t *testing.T) {
		bf.Add('+')
		assert.Equal(t, "++", string(bf.Contents))
	})

	t.Run("test Remove", func(t *testing.T) {
		bf.Remove()
		assert.Equal(t, "+", string(bf.Contents))
	})
}
