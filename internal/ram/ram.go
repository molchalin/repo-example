package ram

import "github.com/molchalin/repo-example/internal/storage"

type Storage struct {
	m map[string]string
}

func New() storage.Storage {
	return &Storage{
		m: make(map[string]string),
	}
}

func (s *Storage) Get(key string) (string, error) {
	if v, ok := s.m[key]; ok {
		return v, nil
	}
	return "", storage.ErrNotFound
}

func (s *Storage) Put(key string, value string) error {
	if _, ok := s.m[key]; ok {
		return storage.ErrAlreadyExists
	}
	s.m[key] = value
	return nil
}
