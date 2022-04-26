package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/molchalin/repo-example/internal/storage"
)

type Server struct {
	storage storage.Storage
}

func New(storage storage.Storage) *Server {
	return &Server{
		storage: storage,
	}
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	val, err := s.storage.Get(id)
	if err == nil {
		fmt.Fprint(w, val)
		return
	}

	status := storageErrToStatus(err)
	w.WriteHeader(status)
}

func (s *Server) Put(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := chi.URLParam(r, "id")

	err = s.storage.Put(id, string(b))
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		return
	}

	status := storageErrToStatus(err)
	w.WriteHeader(status)
}

func storageErrToStatus(err error) int {
	switch err {
	case storage.ErrAlreadyExists:
		return http.StatusConflict
	case storage.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
