package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/molchalin/repo-example/internal/server"
	"github.com/molchalin/repo-example/internal/storage"
)

func main() {
	m := chi.NewRouter()

	st := storage.NewInMemory()

	s := server.New(st)

	m.Get("/{id:[a-z]+}", s.Get)
	m.Put("/{id:[a-z]+}", s.Put)

	http.ListenAndServe(":8080", m)
}
