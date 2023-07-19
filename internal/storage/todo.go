package storage

import "github.com/rhzs/gqlgen-todos/graph/model"

type todo struct {
	values []*model.Todo
}

func NewTodo() *todo {
	return &todo{
		values: []*model.Todo{},
	}
}

func (t *todo) Store(todo *model.Todo) {
	t.values = append(t.values, todo)
}

func (t *todo) Load() []*model.Todo {
	return t.values
}
