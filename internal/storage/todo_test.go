package storage

import (
	"testing"

	"github.com/rhzs/gqlgen-todos/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestTodoStorage(t *testing.T) {
	todo := NewTodo()
	todo.Store(&model.Todo{})
	todoLen := len(todo.Load())

	assert.Equal(t, 1, todoLen)
}
