package query

import (
	"context"
	"testing"

	"github.com/rhzs/gqlgen-todos/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestTodos(t *testing.T) {
	todoStorage := newMockTodoStorage(t)
	todoStorage.On("Load").Return([]*model.Todo{})

	r := NewResolver(todoStorage)

	res, err := r.Todos(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
