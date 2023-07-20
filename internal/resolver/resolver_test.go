package resolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResol(t *testing.T) {
	s := New(nil)

	t.Run("init mutation", func(t *testing.T) {
		assert.NotNil(t, s.Mutation())
	})

	t.Run("init query", func(t *testing.T) {
		assert.NotNil(t, s.Query())
	})
}
