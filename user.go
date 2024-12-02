package fleet

import (
	"net/http"
	"strconv"
	"time"
)

type (
	User struct {
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		ID                 int       `json:"id"`
		Name               string    `json:"name"`
		Email              string    `json:"email"`
		Enabled            bool      `json:"enabled"`
		ForcePasswordReset bool      `json:"force_password_reset"`
		GravatarURL        string    `json:"gravatar_url"`
		SSOEnabled         bool      `json:"sso_enabled"`
		GlobalRole         string    `json:"global_role"`
		Teams              []*Team   `json:"teams"`
	}

	UserNew struct {
		Email                string  `json:"email"`
		InviteToken          string  `json:"invite_token"`
		Name                 string  `json:"name"`
		Password             string  `json:"password"`
		PasswordConfirmation string  `json:"password_confirmation"`
		Teams                []*Team `json:"teams"`
	}

	Team struct {
		ID          int       `json:"id"`
		CreatedAt   time.Time `json:"created_at"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Role        string    `json:"role"`
	}
)

// CreateUser creates the given user, and returns a formatted User object. Use CreateUserWithoutInvite if you don't want
// to send the user an invitation.
func (s *Service) CreateUser(user *UserNew) (*User, error) {
	resp := struct {
		User *User `json:"user"`
	}{}

	if _, err := s.makeRequest(http.MethodPost, "users", user, &resp); err != nil {
		return nil, err
	}

	return resp.User, nil
}

// CreateUserWithoutInvite creates the given user, and returns a formatted User object.
func (s *Service) CreateUserWithoutInvite(user *UserNew) (*User, error) {
	user.InviteToken = ""

	resp := struct {
		User *User `json:"user"`
	}{}

	if _, err := s.makeRequest(http.MethodPost, "users/admin", user, &resp); err != nil {
		return nil, err
	}

	return resp.User, nil
}

// DeleteUser deletes the user with the corresponding ID.
func (s *Service) DeleteUser(id int) error {
	if _, err := s.makeRequest(http.MethodDelete, "users/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

// FindUserByID returns the user with the corresponding ID.
func (s *Service) FindUserByID(id int) (*User, error) {
	resp := struct {
		User *User `json:"user"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "users/"+strconv.Itoa(id), nil, &resp); err != nil {
		return nil, err
	}

	return resp.User, nil
}

// FindUsers returns all users from the Fleet API. The query field can be used to dynamically match user names or
// emails.
func (s *Service) FindUsers(query string) ([]*User, error) {
	resp := struct {
		Users []*User `json:"users"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "users?query="+query, nil, &resp); err != nil {
		return nil, err
	}

	return resp.Users, nil
}

func (s *Service) FindMe() (*User, error) {
	resp := struct {
		User *User `json:"user"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "me", nil, &resp); err != nil {
		return nil, err
	}

	return resp.User, nil
}

// RequirePasswordReset logs the given user out of all active sessions, and forces them to reset their password.
func (s *Service) RequirePasswordReset(id int) error {
	if _, err := s.makeRequest(http.MethodPost, "users/"+strconv.Itoa(id)+"/require_password_reset", nil, nil); err != nil {
		return err
	}

	return nil
}
