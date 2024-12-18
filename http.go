package fleet

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type (
	APIError struct {
		Message string `json:"message"`
		Errors  []struct {
			Name   string `json:"name"`
			Reason string `json:"reason"`
		} `json:"errors"`
		Uuid string `json:"uuid"`
	}

	httpDebug struct {
		Body     string
		Code     int
		Endpoint string
		Start    time.Time
	}
)

func (err *APIError) String() string {
	errs := make([]string, len(err.Errors))

	for i, e := range err.Errors {
		errs[i] = e.Name + ": " + e.Reason
	}

	return strings.Join(errs, "; ")
}

// execRequest is the underlying function for HTTP requests. It handles debugging statements, and simple error handling
func (s *Service) execRequest(req *http.Request) ([]byte, error) {
	debug := httpDebug{
		Endpoint: req.URL.Path,
		Start:    time.Now(),
	}
	defer s.printDebugHTTP(&debug)

	resp, err := s.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if s.debug {
		debug.Code = resp.StatusCode
	}

	var responseBytes []byte

	if resp.Body != nil {
		responseBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}

		if s.debug {
			// if len(responseBytes) > 32768 {
			// 	debug.Body = "body too long for debug: " + cast.ToString(len(responseBytes))
			// } else {
			debug.Body = string(responseBytes)
			// }
		}
	}

	if resp.StatusCode/100 != 2 {
		return responseBytes, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return responseBytes, nil
}

// makeRequest is called by internal functions only. It handles the common HTTP request requirements for the API.
func (s *Service) makeRequest(method string, endpoint string, body any, object any) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(strings.ToUpper(method), s.url+"/"+endpoint, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if s.key != "" {
		req.Header.Set("Authorization", "Bearer "+s.key)
	}

	resp, err := s.execRequest(req)
	if err != nil {
		var apiError APIError

		jsonErr := json.Unmarshal(resp, &apiError)
		if jsonErr != nil {
			return nil, fmt.Errorf("error making request: %w", err)
		}

		return nil, errors.New(apiError.String())
	}

	if resp != nil && object != nil {
		if err = json.Unmarshal(resp, &object); err != nil {
			return nil, fmt.Errorf("error unmarshalling response: %w", err)
		}
	}

	return resp, nil
}

func (s *Service) printDebugHTTP(debug *httpDebug) {
	if s.debug {
		fmt.Printf("\nENDPOINT: %v\nSTATUS CODE: %v\nTIME STARTED: %v\nTIME ENDED: %v\nTIME TAKEN: %v\nBODY: %v\n", debug.Endpoint, debug.Code, debug.Start, time.Now(), time.Since(debug.Start), debug.Body)
	}
}
