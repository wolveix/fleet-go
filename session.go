package fleet

import (
	"net/http"
	"strconv"
	"time"
)

type Session struct {
	ID        int       `json:"session_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

// DeleteSession deletes the session with the corresponding ID.
func (s *Service) DeleteSession(id int) error {
	if _, err := s.makeRequest(http.MethodDelete, "sessions/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

// FindSessionByID returns the session with the corresponding ID.
func (s *Service) FindSessionByID(id int) (*Session, error) {
	var resp Session

	if _, err := s.makeRequest(http.MethodGet, "sessions/"+strconv.Itoa(id), nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// FindSessionsByUserID returns all sessions for the given user ID.
func (s *Service) FindSessionsByUserID(id int) ([]*Session, error) {
	resp := struct {
		Sessions []*Session `json:"sessions"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "users/"+strconv.Itoa(id)+"/sessions", nil, &resp); err != nil {
		return nil, err
	}

	return resp.Sessions, nil
}
