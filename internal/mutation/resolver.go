package mutation

import (
	"github.com/rhzs/gqlgen-todos/graph"
	"github.com/rhzs/gqlgen-todos/graph/model"
)

type resolver struct {
	todoStorage todoStorage
}

var _ graph.MutationResolver = &resolver{}

func NewResolver(todoStorage todoStorage) *resolver {
	return &resolver{
		todoStorage: todoStorage,
	}
}

//go:generate mockery --name=todoStorage --case=underscore --inpackage --testonly
type todoStorage interface {
	Store(todo *model.Todo)
}
