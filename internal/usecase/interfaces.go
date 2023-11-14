package usecase

import (
	"context"
	"todo/internal/entity"
)

type (
	Todo interface {
		List(context.Context) ([]entity.Todo, error)
		Add(context.Context, entity.Todo) (entity.Todo, error)
	}
)
