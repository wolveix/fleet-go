package fleet

import (
	"net/http"
	"strconv"
	"time"
)

type Host struct {
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt                 time.Time   `json:"updated_at"`
	ID                        int         `json:"id"`
	DetailUpdatedAt           time.Time   `json:"detail_updated_at"`
	LastRestartedAt           time.Time   `json:"last_restarted_at"`
	SoftwareUpdatedAt         time.Time   `json:"software_updated_at"`
	LabelUpdatedAt            time.Time   `json:"label_updated_at"`
	PolicyUpdatedAt           time.Time   `json:"policy_updated_at"`
	LastEnrolledAt            time.Time   `json:"last_enrolled_at"`
	SeenTime                  time.Time   `json:"seen_time"`
	Hostname                  string      `json:"hostname"`
	UUID                      string      `json:"uuid"`
	Platform                  string      `json:"platform"`
	OsqueryVersion            string      `json:"osquery_version"`
	OSVersion                 string      `json:"os_version"`
	Build                     string      `json:"build"`
	PlatformLike              string      `json:"platform_like"`
	CodeName                  string      `json:"code_name"`
	Uptime                    int64       `json:"uptime"`
	Memory                    int         `json:"memory"`
	CPUType                   string      `json:"cpu_type"`
	CPUSubtype                string      `json:"cpu_subtype"`
	CPUBrand                  string      `json:"cpu_brand"`
	CPUPhysicalCores          int         `json:"cpu_physical_cores"`
	CPULogicalCores           int         `json:"cpu_logical_cores"`
	HardwareVendor            string      `json:"hardware_vendor"`
	HardwareModel             string      `json:"hardware_model"`
	HardwareVersion           string      `json:"hardware_version"`
	HardwareSerial            string      `json:"hardware_serial"`
	ComputerName              string      `json:"computer_name"`
	DisplayName               string      `json:"display_name"`
	PublicIP                  string      `json:"public_ip"`
	PrimaryIP                 string      `json:"primary_ip"`
	PrimaryMac                string      `json:"primary_mac"`
	DistributedInterval       int         `json:"distributed_interval"`
	ConfigTLSRefresh          int         `json:"config_tls_refresh"`
	LoggerTLSPeriod           int         `json:"logger_tls_period"`
	Additional                struct{}    `json:"additional"`
	Status                    string      `json:"status"`
	DisplayText               string      `json:"display_text"`
	TeamID                    interface{} `json:"team_id"`
	TeamName                  interface{} `json:"team_name"`
	GigsDiskSpaceAvailable    float64     `json:"gigs_disk_space_available"`
	PercentDiskSpaceAvailable int         `json:"percent_disk_space_available"`
	GigsTotalDiskSpace        float64     `json:"gigs_total_disk_space"`
	PackStats                 []struct {
		PackID     int    `json:"pack_id"`
		PackName   string `json:"pack_name"`
		Type       string `json:"type"`
		QueryStats []struct {
			ScheduledQueryName string      `json:"scheduled_query_name"`
			ScheduledQueryID   int         `json:"scheduled_query_id"`
			QueryName          string      `json:"query_name"`
			DiscardData        bool        `json:"discard_data"`
			LastFetched        interface{} `json:"last_fetched"`
			AutomationsEnabled bool        `json:"automations_enabled"`
			Description        string      `json:"description"`
			PackName           string      `json:"pack_name"`
			AverageMemory      int         `json:"average_memory"`
			DenyListed         bool        `json:"denylisted"`
			Executions         int         `json:"executions"`
			Interval           int         `json:"interval"`
			LastExecuted       time.Time   `json:"last_executed"`
			OutputSize         int         `json:"output_size"`
			SystemTime         int         `json:"system_time"`
			UserTime           int         `json:"user_time"`
			WallTime           int         `json:"wall_time"`
		} `json:"query_stats"`
	} `json:"pack_stats"`
	Issues struct {
		FailingPoliciesCount         int `json:"failing_policies_count"`
		CriticalVulnerabilitiesCount int `json:"critical_vulnerabilities_count"`
		TotalIssuesCount             int `json:"total_issues_count"`
	} `json:"issues"`
	Geolocation struct {
		CountryISO string `json:"country_iso"`
		CityName   string `json:"city_name"`
		Geometry   struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"geolocation"`
	MDM struct {
		EncryptionKeyAvailable bool   `json:"encryption_key_available"`
		EnrollmentStatus       string `json:"enrollment_status"`
		DepProfileError        bool   `json:"dep_profile_error"`
		Name                   string `json:"name"`
		ServerUrl              string `json:"server_url"`
	} `json:"mdm"`
	Software []struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		Version         string `json:"version"`
		Source          string `json:"source"`
		GeneratedCpe    string `json:"generated_cpe"`
		Vulnerabilities []struct {
			CVE               string  `json:"cve"`
			DetailsLink       string  `json:"details_link"`
			CVSSScore         float64 `json:"cvss_score"`
			EPSSProbability   float64 `json:"epss_probability"`
			CISAKnownExploit  bool    `json:"cisa_known_exploit"`
			CVEPublished      string  `json:"cve_published"`
			CVEDescription    string  `json:"cve_description"`
			ResolvedInVersion string  `json:"resolved_in_version"`
		} `json:"vulnerabilities"`
		InstalledPaths []string `json:"installed_paths"`
	} `json:"software"`
	Policies []Policy `json:"policies"`
}

// DeleteHost deletes the host with the corresponding ID.
func (s *Service) DeleteHost(id int) error {
	if _, err := s.makeRequest(http.MethodDelete, "hosts/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

// FindHostByID returns the host with the corresponding ID.
func (s *Service) FindHostByID(id int) (*Host, error) {
	resp := struct {
		Host *Host `json:"host"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "hosts/"+strconv.Itoa(id), nil, &resp); err != nil {
		return nil, err
	}

	return resp.Host, nil
}

// FindHosts returns all hosts from the Fleet API.
func (s *Service) FindHosts() ([]*Host, error) {
	resp := struct {
		Hosts []*Host `json:"hosts"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "hosts", nil, &resp); err != nil {
		return nil, err
	}

	return resp.Hosts, nil
}

// FindHostsByLabel returns all hosts from the Fleet API within the given label.
func (s *Service) FindHostsByLabel(labelID int) ([]*Host, error) {
	resp := struct {
		Hosts []*Host `json:"hosts"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "labels/"+strconv.Itoa(labelID)+"/hosts", nil, &resp); err != nil {
		return nil, err
	}

	return resp.Hosts, nil
}

// RefetchHost flags that the host details, labels, and policies should be refetched the next time the host checks in.
func (s *Service) RefetchHost(id int) error {
	if _, err := s.makeRequest(http.MethodPost, "hosts/"+strconv.Itoa(id)+"/refetch", nil, nil); err != nil {
		return err
	}

	return nil
}

type Activity struct {
	CreatedAt     time.Time `json:"created_at"`
	ActorID       int       `json:"actor_id"`
	ActorFullName string    `json:"actor_full_name"`
	ID            int       `json:"id"`
	ActorGravatar string    `json:"actor_gravatar"`
	ActorEmail    string    `json:"actor_email"`
	Type          string    `json:"type"`
	Details       struct {
		HostID            int    `json:"host_id"`
		HostDisplayName   string `json:"host_display_name"`
		SoftwareTitle     string `json:"software_title"`
		ScriptExecutionId string `json:"script_execution_id"`
		Status            string `json:"status"`
	} `json:"details"`
}

// FindHostPastActivity returns all past activities for the corresponding host.
func (s *Service) FindHostPastActivity(id int) ([]*Activity, error) {
	resp := struct {
		Activities []*Activity `json:"activities"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "hosts/"+strconv.Itoa(id)+"/activities", nil, &resp); err != nil {
		return nil, err
	}

	return resp.Activities, nil
}

// FindHostUpcomingActivity returns all upcoming activities for the corresponding host.
func (s *Service) FindHostUpcomingActivity(id int) ([]*Activity, error) {
	resp := struct {
		Activities []*Activity `json:"activities"`
	}{}

	if _, err := s.makeRequest(http.MethodGet, "hosts/"+strconv.Itoa(id)+"/activities/upcoming", nil, &resp); err != nil {
		return nil, err
	}

	return resp.Activities, nil
}
