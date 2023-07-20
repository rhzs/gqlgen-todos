package resolver

import (
	"github.com/rhzs/gqlgen-todos/graph"
	"github.com/rhzs/gqlgen-todos/graph/model"
	"github.com/rhzs/gqlgen-todos/internal/mutation"
	"github.com/rhzs/gqlgen-todos/internal/query"
)

type resolver struct {
	todoStorage todoStorage
}

func New(todoStorage todoStorage) *resolver {
	return &resolver{
		todoStorage: todoStorage,
	}
}

// Mutation returns MutationResolver implementation.
func (r *resolver) Mutation() graph.MutationResolver {
	return mutation.NewResolver(r.todoStorage)
}

// Query returns QueryResolver implementation.
func (r *resolver) Query() graph.QueryResolver {
	return query.NewResolver(r.todoStorage)
}

//go:generate mockery --name=todoStorage --case=underscore --inpackage --testonly
type todoStorage interface {
	Store(todo *model.Todo)
	Load() []*model.Todo
}
