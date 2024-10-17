package logstore

import "fmt"

type Store interface {
	// Add a string key and any value to the store
	Add(key string, val any) error
}

type logStore struct {
	path string
}

func (s *logStore) Add(key string, val any) error {
	fmt.Printf("Adding %v key and %v value\n", key, val)
	return nil
}

func NewStore(storePath string) (Store, error) {
	return &logStore{}, nil
}
