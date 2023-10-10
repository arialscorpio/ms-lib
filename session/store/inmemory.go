package store

import (
	"fmt"

	"github.com/google/uuid"

	"my.gopherworld.dev/session"
)

type InMemSessionStore struct {
	sessions map[string]session.Session
}

func (ss *InMemSessionStore) Get(id string) (*session.Session, error) {
	if sess, ok := ss.sessions[id]; ok {
		return &sess, nil
	}
	return nil, session.ErrSessionNotFound
}

func (ss *InMemSessionStore) Create(s session.Session) (string, error) {
	randID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("session-ID generate error, %w", err)
	}

	id := randID.String()
	ss.sessions[id] = s
	return id, nil
}

func (ss *InMemSessionStore) Update(id string, s session.Session) {
	ss.sessions[id] = s
}

func (ss *InMemSessionStore) Delete(id string) {
	delete(ss.sessions, id)
}
