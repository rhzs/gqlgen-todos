package mutation

import (
	"context"
	"testing"

	"github.com/rhzs/gqlgen-todos/graph/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTodo(t *testing.T) {
	todoStorage := newMockTodoStorage(t)
	todoStorage.On("Store", mock.Anything)

	r := NewResolver(todoStorage)

	res, err := r.CreateTodo(context.Background(), model.NewTodo{})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
