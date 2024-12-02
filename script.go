package fleet

import (
	"net/http"
	"strconv"
	"time"
)

type Script struct {
	ID        int         `json:"id"`
	TeamID    interface{} `json:"team_id"`
	Name      string      `json:"name"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// DeleteScript deletes the script with the corresponding ID.
func (s *Service) DeleteScript(id int) error {
	if _, err := s.makeRequest(http.MethodDelete, "scripts/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

// FindScriptByID returns the script with the corresponding ID.
func (s *Service) FindScriptByID(id int) (*Script, error) {
	resp := struct {
		Script *Script `json:"script"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "scripts/"+strconv.Itoa(id), nil, &resp); err != nil {
		return nil, err
	}

	return resp.Script, nil
}

// FindScriptsByHostID returns script metadata for the given host.
func (s *Service) FindScriptsByHostID(hostID int) ([]*Script, error) {
	resp := struct {
		Scripts []*Script `json:"scripts"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "hosts/"+strconv.Itoa(hostID)+"/scripts", nil, &resp); err != nil {
		return nil, err
	}

	return resp.Scripts, nil
}

// FindScripts returns all scripts from the Fleet API.
func (s *Service) FindScripts() ([]*Script, error) {
	resp := struct {
		Scripts []*Script `json:"scripts"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "scripts", nil, &resp); err != nil {
		return nil, err
	}

	return resp.Scripts, nil
}

// RunScript runs the given raw script on the given host, and returns the UUID of the script execution.
func (s *Service) RunScript(hostID int, scriptContents string) (string, error) {
	body := struct {
		HostID         int    `json:"host_id"`
		ScriptContents string `json:"script_contents"`
	}{
		hostID, scriptContents,
	}

	resp := struct {
		ExecutionID string `json:"execution_id"`
	}{}

	if _, err := s.makeRequest(http.MethodPost, "scripts/run", body, &resp); err != nil {
		return "", err
	}

	return resp.ExecutionID, nil
}

// RunScriptID runs the given script on the given host, and returns the UUID of the script execution.
func (s *Service) RunScriptID(hostID int, scriptID int) (string, error) {
	body := struct {
		HostID   int `json:"host_id"`
		ScriptID int `json:"script_id"`
	}{
		hostID, scriptID,
	}

	resp := struct {
		ExecutionID string `json:"execution_id"`
	}{}

	if _, err := s.makeRequest(http.MethodPost, "scripts/run", body, &resp); err != nil {
		return "", err
	}

	return resp.ExecutionID, nil
}

type ScriptResult struct {
	ScriptContents string    `json:"script_contents"`
	ExitCode       int       `json:"exit_code"`
	Output         string    `json:"output"`
	Message        string    `json:"message"`
	Hostname       string    `json:"hostname"`
	HostTimeout    bool      `json:"host_timeout"`
	HostID         int       `json:"host_id"`
	ID             string    `json:"execution_id"`
	Runtime        int       `json:"runtime"`
	CreatedAt      time.Time `json:"created_at"`
}

// FindScriptResult returns the script result with the corresponding ID.
func (s *Service) FindScriptResult(id int) (*ScriptResult, error) {
	scriptResult := ScriptResult{}

	if _, err := s.makeRequest(http.MethodGet, "scripts/"+strconv.Itoa(id), nil, &scriptResult); err != nil {
		return nil, err
	}

	return &scriptResult, nil
}
