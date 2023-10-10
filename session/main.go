package session

import (
	"fmt"

	"github.com/google/uuid"
)

type SessionStore interface {
	Get(id string) *Session
	Create(s Session) string
	Update(id string, s Session)
	Delete(id string)
}

type InMemSessionStore struct {
	sessions map[string]Session
}

func (ss *InMemSessionStore) Get(id string) *Session {
	if sess, ok := ss.sessions[id]; ok {
		return &sess
	}
	return nil
}

func (ss *InMemSessionStore) Create(s Session) (string, error) {
	randID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("session-ID generate error, %w", err)
	}

	id := randID.String()
	ss.sessions[id] = s
	return id, nil
}

func (ss *InMemSessionStore) Update(id string, s Session) {
	ss.sessions[id] = s
}

func (ss *InMemSessionStore) Delete(id string) {
	delete(ss.sessions, id)
}
