package usecase

import (
	"context"
	"fmt"
	"todo/internal/entity"
	"todo/internal/repository"
)

type TodoUseCase struct {
}

func New() *TodoUseCase {
	return &TodoUseCase{}
}

func (uc *TodoUseCase) List(ctx context.Context) ([]entity.Todo, error) {
	r := &repository.TodoRepository{}
	todo, err := r.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error")
	}
	return todo, nil
}

func (uc *TodoUseCase) Add(ctx context.Context, t entity.Todo) (entity.Todo, error) {
	r := &repository.TodoRepository{}
	todo, err := r.Add(ctx, t)
	if err != nil {
		return entity.Todo{}, fmt.Errorf("Error")
	}
	return todo, nil
}
