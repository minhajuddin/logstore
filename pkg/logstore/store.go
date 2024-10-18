package logstore

import "fmt"

type LogStore interface {
	// Add a string key and any value to the store
	Add(key string, val string) error
	Get(key string) (string, error)
}

type logStore struct {
	path string
}

func (s *logStore) Add(key string, val string) error {
	fmt.Printf("Adding %v key and %v value\n", key, val)
	return nil
}

func (s *logStore) Get(key string) (string, error) {
	return "banana", nil
}

// NewStore creates a new store at the given path or loads an existing store
func NewStore(storePath string) (LogStore, error) {
	return &logStore{}, nil
}
