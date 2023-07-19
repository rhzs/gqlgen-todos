package query

import (
	"context"

	"github.com/rhzs/gqlgen-todos/graph"
	"github.com/rhzs/gqlgen-todos/graph/model"
)

type resolver struct {
	todoStorage todoStorage
}

var _ graph.QueryResolver = &resolver{}

func NewResolver(todoStorage todoStorage) *resolver {
	return &resolver{
		todoStorage: todoStorage,
	}
}

func (r *resolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todoStorage.Load(), nil
}

type todoStorage interface {
	Load() []*model.Todo
}
