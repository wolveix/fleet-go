package fleet

type (
	Software struct {
		ID              int             `json:"id"`
		Name            string          `json:"name"`
		Version         string          `json:"version"`
		Source          string          `json:"source"`
		GeneratedCpe    string          `json:"generated_cpe"`
		Vulnerabilities []Vulnerability `json:"vulnerabilities"`
		InstalledPaths  []string        `json:"installed_paths"`
	}

	Vulnerability struct {
		CVE               string  `json:"cve"`
		DetailsLink       string  `json:"details_link"`
		CVSSScore         float64 `json:"cvss_score"`
		EPSSProbability   float64 `json:"epss_probability"`
		CISAKnownExploit  bool    `json:"cisa_known_exploit"`
		CVEPublished      string  `json:"cve_published"`
		CVEDescription    string  `json:"cve_description"`
		ResolvedInVersion string  `json:"resolved_in_version"`
	}
)
