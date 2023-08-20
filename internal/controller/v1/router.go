package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"todo/internal/usecase"
)

func NewRouter(handler *chi.Mux, t usecase.Todo) {
	handler.Use(middleware.Logger)

	handler.Route("/v1", func(handler chi.Router) {
		newTodoRoutes(handler, t)
	})
}
