package logstore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	s, err := NewStore("/tmp/test.logstore.db")
	assert.Nil(t, err)
	assert.NotNil(t, s)

	assert.Nil(t, s.Add("a", "test"))
}

func TestGet(t *testing.T) {
	s, err := NewStore("/tmp/test.logstore.db")
	assert.Nil(t, err)
	assert.Nil(t, s.Add("boy", "Danish"))
	v, err := s.Get("boy")

	assert.Nil(t, err)
	assert.Equal(t, v, "Danish")
}
