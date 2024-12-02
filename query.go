package fleet

import (
	"net/http"
	"strconv"
	"time"
)

type (
	Query struct {
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		ID                 int       `json:"id"`
		Name               string    `json:"name"`
		Description        string    `json:"description"`
		Query              string    `json:"query"`
		TeamID             int       `json:"team_id"`
		Interval           int       `json:"interval"`
		Platform           string    `json:"platform"`
		Version            string    `json:"version"`
		AutomationsEnabled bool      `json:"automations_enabled"`
		Logging            string    `json:"logging"`
		Saved              bool      `json:"saved"`
		ObserverCanRun     bool      `json:"observer_can_run"`
		DiscardData        bool      `json:"discard_data"`
		AuthorID           int       `json:"author_id"`
		AuthorName         string    `json:"author_name"`
		AuthorEmail        string    `json:"author_email"`
		Packs              []struct {
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
			ID          int       `json:"id"`
			Name        string    `json:"name"`
			Description string    `json:"description"`
			Platform    string    `json:"platform"`
			Disabled    bool      `json:"disabled"`
		} `json:"packs"`
		Stats struct {
			SystemTimeP50   float64 `json:"system_time_p50"`
			SystemTimeP95   float64 `json:"system_time_p95"`
			UserTimeP50     float64 `json:"user_time_p50"`
			UserTimeP95     float64 `json:"user_time_p95"`
			TotalExecutions float64 `json:"total_executions"`
		} `json:"stats"`
	}

	QueryResult struct {
		QueryID            int `json:"query_id"`
		TargetedHostCount  int `json:"targeted_host_count"`
		RespondedHostCount int `json:"responded_host_count"`
		Results            []struct {
			HostID int `json:"host_id"`
			Rows   []struct {
				BuildDistro   string `json:"build_distro"`
				BuildPlatform string `json:"build_platform"`
				ConfigHash    string `json:"config_hash"`
				ConfigValid   string `json:"config_valid"`
				Extensions    string `json:"extensions"`
				InstanceID    string `json:"instance_id"`
				PID           string `json:"pid"`
				PlatformMask  string `json:"platform_mask"`
				StartTime     string `json:"start_time"`
				UUID          string `json:"uuid"`
				Version       string `json:"version"`
				Watcher       string `json:"watcher"`
			} `json:"rows"`
			Error string `json:"error"`
		} `json:"results"`
	}
)

// FindQueries returns all queries from the Fleet API. The query field can be used to dynamically match query names.
func (s *Service) FindQueries(query string) ([]*Query, error) {
	resp := struct {
		Queries []*Query `json:"queries"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "queries?query="+query, nil, &resp); err != nil {
		return nil, err
	}

	return resp.Queries, nil
}

// RunQueryOnHosts executes the given query on the given hosts.
func (s *Service) RunQueryOnHosts(id int, hostIDs ...int) (*QueryResult, error) {
	queryResult := QueryResult{}

	req := struct {
		HostIDs []int `json:"host_ids"`
	}{HostIDs: hostIDs}

	if _, err := s.makeRequest(http.MethodPost, "queries/"+strconv.Itoa(id)+"/run", req, &queryResult); err != nil {
		return nil, err
	}

	return &queryResult, nil
}
