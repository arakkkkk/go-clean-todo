package v1

import (
	"encoding/json"
	"net/http"
	"time"
	"todo/internal/entity"
	"todo/internal/usecase"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/render"
)

type todoRoutes struct {
	t usecase.Todo
}

func newTodoRoutes(handler chi.Router, t usecase.Todo) {
	handler.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r := &todoRoutes{t}
	handler.Get("/list", r.list)
	handler.Post("/add", r.add)
}

func (router *todoRoutes) list(w http.ResponseWriter, r *http.Request) {
	todos, err := router.t.List(r.Context())
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	}
	response, err := json.MarshalIndent(todos, "", "")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

type addRequest struct {
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Priority  int       `json:"priority"`
	CreatedAt time.Time `json:"create_at"`
}

func (a *addRequest) Bind(r *http.Request) error {
	return nil
}

func (router *todoRoutes) add(w http.ResponseWriter, r *http.Request) {
	request := &addRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	}
	todo, err := router.t.Add(r.Context(), entity.Todo{
		Title:    request.Title,
		Priority: request.Priority,
	})
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	}
	response, err := json.MarshalIndent(todo, "", "")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
