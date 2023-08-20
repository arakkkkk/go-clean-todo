package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"todo/internal/controller/v1"
	"todo/internal/usecase"
)

func Run() {
	r := chi.NewRouter()
	t := usecase.New()
	v1.NewRouter(r, t)

	http.ListenAndServe(":8080", r)
}
