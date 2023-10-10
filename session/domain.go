package session

import (
	"errors"
	"time"
)

var ErrKeyNotFound = errors.New("key not found in session")

type Session struct {
	data   map[string]string
	expiry time.Time
}

func (s *Session) Expired() bool {
	return !time.Now().Before(s.expiry)
}

func (s *Session) Get(k string) (string, error) {
	if v, ok := s.data[k]; ok {
		return v, nil
	}
	return "", ErrKeyNotFound
}

func (s *Session) Set(k, v string) {
	s.data[k] = v
}

func (s *Session) Delete(k string) {
	delete(s.data, k)
}

// Store defines a storage to store user sessions.
type Store interface {
	Get(id string) (*Session, error)
	Create(s Session) string
	Update(id string, s Session)
	Delete(id string)
}

var ErrSessionNotFound = errors.New("session not found")
