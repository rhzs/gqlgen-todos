package mutation

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/rhzs/gqlgen-todos/graph/model"
)

func (r *resolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", randInt),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todoStorage.Store(todo)

	return todo, nil
}
