package fleet

import (
	"net/http"
)

// ChangePassword allows you to change your own password.
func (s *Service) ChangePassword(oldPassword string, newPassword string) error {
	if _, err := s.makeRequest(http.MethodPost, "change_password",
		map[string]string{"old_password": oldPassword, "new_password": newPassword}, nil); err != nil {
		return err
	}

	return nil
}

// Login will create a new session with the Fleet API, and will return the auth token alongside your user details.
func (s *Service) Login(email string, password string) (string, *User, error) {
	resp := struct {
		User  *User  `json:"user"`
		Token string `json:"token"`
	}{}

	if _, err := s.makeRequest(http.MethodPost, "login",
		map[string]string{"email": email, "password": password}, &resp); err != nil {
		return "", nil, err
	}

	return resp.Token, resp.User, nil
}

// Logout will close the currently authenticated session, invalidating the in-use token/API key.
func (s *Service) Logout() error {
	if _, err := s.makeRequest(http.MethodPost, "logout", nil, nil); err != nil {
		return err
	}

	return nil
}

// SendForgottenPasswordEmail will send a password reset email to the given email address, assuming that a user exists
// for it.
func (s *Service) SendForgottenPasswordEmail(email string) error {
	if _, err := s.makeRequest(http.MethodPost, "forgot_password",
		map[string]string{"email": email}, nil); err != nil {
		return err
	}

	return nil
}
