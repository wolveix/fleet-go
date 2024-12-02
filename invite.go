package fleet

import (
	"net/http"
	"strconv"
)

// DeleteInvite deletes the invite with the corresponding ID.
func (s *Service) DeleteInvite(id int) error {
	if _, err := s.makeRequest(http.MethodDelete, "invites/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

// FindInvites returns all invites from the Fleet API. The query field can be used to dynamically match invite names or
// emails.
func (s *Service) FindInvites(query string) ([]*User, error) {
	resp := struct {
		Invites []*User `json:"invites"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "invites?query="+query, nil, &resp); err != nil {
		return nil, err
	}

	return resp.Invites, nil
}

// VerifyInvite verifies the invite with the corresponding ID.
func (s *Service) VerifyInvite(token string) (*User, error) {
	resp := struct {
		User *User `json:"invite"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "invites"+token, nil, &resp); err != nil {
		return nil, err
	}

	return resp.User, nil
}
