package repository

import (
	"context"
	"log"
	"time"
	"todo/ent"
	"todo/internal/entity"
	entPkg "todo/pkg/ent"
)

type TodoRepository struct{}

func entTodoBindEntiryTodo(t *ent.Todo) entity.Todo {
	entityTodo := entity.Todo{
		Title:     t.Title,
		Completed: t.Completed,
		Priority:  t.Priority,
		CretedAt:  t.CreatedAt,
	}
	return entityTodo
}
func entTodosBindEntiryTodos(t []*ent.Todo) []entity.Todo {
	entityTodos := []entity.Todo{}
	for _, v := range t {
		entityTodos = append(entityTodos, entTodoBindEntiryTodo(v))
	}
	return entityTodos[:]

}

func (r *TodoRepository) List(ctx context.Context) ([]entity.Todo, error) {
	client, err := entPkg.Open()
	if err != nil {
		log.Fatalf("failed connecting db: %v", err)
	}
	defer client.Close()
	todos, err := client.Todo.Query().All(ctx)
	entityTodos := entTodosBindEntiryTodos(todos)

	if err != nil {
		log.Fatalf("failed query user: %v", err)
	}

	return entityTodos, nil
}

func (r *TodoRepository) Add(ctx context.Context, t entity.Todo) (entity.Todo, error) {
	client, err := entPkg.Open()
	if err != nil {
		log.Fatalf("failed connecting db: %v", err)
	}
	defer client.Close()
	todo, err := client.Todo.
		Create().
		SetTitle(t.Title).
		SetCompleted(t.Completed).
		SetPriority(t.Priority).
		SetCreatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating todo: %v", err)
	}

	entityTodo := entTodoBindEntiryTodo(todo)
	return entityTodo, nil
}
