package session

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	// "my.gopherworld.dev/session"
)

const defaultIdentKey = "sid"

var (
	ErrNoSession      = errors.New("no exisiting session")
	ErrSessionExpired = errors.New("session expired")
)

type SessionController struct {
	// identKey defines the name of the http session cookie.
	identKey string
	store    Store
}

func NewController(idenKey string, store Store) *SessionController {
	return &SessionController{idenKey, store}
}

func (sc *SessionController) Initialize(r *http.Request) (*Session, error) {
	c, err := r.Cookie(sc.identKey)
	if err == http.ErrNoCookie {
		return nil, ErrNoSession
	}

	sid := c.Value
	sess, err := sc.store.Get(sid)
	if err == ErrSessionNotFound {
		return nil, ErrNoSession
	}
	if err != nil {
		return nil, fmt.Errorf("session get error: %w", err)
	}
	if sess.Expired() {
		return nil, ErrSessionExpired
	}

	return sess, nil
}

func (sc *SessionController) WriteToResponse(id string, s Session, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "sid",
		Value:   id,
		Path:    "/",
		Expires: time.Now().Add(15 * time.Minute),
		Secure:  true,
	})
}

func (sc *SessionController) Persist(id string, s Session) {
	sc.store.Update(id, s)
}
