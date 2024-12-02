package fleet

import (
	"net/http"
	"strings"
	"time"
)

type Service struct {
	key string
	url string

	debug bool
	http  *http.Client
}

func New(url string, key string, timeout time.Duration, debug bool) *Service {
	if !strings.HasSuffix(url, "/api/v1/fleet") {
		url = url + "/api/v1/fleet"
	}

	return &Service{
		key:   key,
		url:   url,
		debug: debug,
		http:  &http.Client{Timeout: timeout},
	}
}

func (s *Service) SetKey(key string) {
	s.key = key
}

type Version struct {
	Version   string `json:"version"`
	Branch    string `json:"branch"`
	Revision  string `json:"revision"`
	GoVersion string `json:"go_version"`
	BuildDate string `json:"build_date"`
	BuildUser string `json:"build_user"`
}

func (s *Service) FindVersion() (*Version, error) {
	resp := Version{}

	if _, err := s.makeRequest(http.MethodGet, "version", nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
