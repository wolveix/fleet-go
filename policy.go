package fleet

import (
	"net/http"
	"time"
)

type Policy struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Query              string    `json:"query"`
	Description        string    `json:"description"`
	Critical           bool      `json:"critical"`
	AuthorID           int       `json:"author_id"`
	AuthorName         string    `json:"author_name"`
	AuthorEmail        string    `json:"author_email"`
	TeamID             string    `json:"team_id"`
	Resolution         string    `json:"resolution"`
	Platform           string    `json:"platform"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	PassingHostCount   int       `json:"passing_host_count"`
	FailingHostCount   int       `json:"failing_host_count"`
	HostCountUpdatedAt time.Time `json:"host_count_updated_at"`
}

// DeletePolicy deletes the policy with the corresponding ID.
func (s *Service) DeletePolicy(id int) error {
	req := struct {
		IDs []int `json:"ids"`
	}{
		[]int{id},
	}

	if _, err := s.makeRequest(http.MethodDelete, "global/policies/delete", req, nil); err != nil {
		return err
	}

	return nil
}

// FindPolicies returns all policies from the Fleet API.
func (s *Service) FindPolicies() ([]*Policy, error) {
	resp := struct {
		Policies []*Policy `json:"policies"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "global/policies", nil, &resp); err != nil {
		return nil, err
	}

	return resp.Policies, nil
}
