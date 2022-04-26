package storage

import "errors"

var (
	ErrNotFound      = errors.New("Not Found")
	ErrAlreadyExists = errors.New("Already Exists")
)

type Storage interface {
	Get(key string) (string, error)
	Put(key, value string) error
}
