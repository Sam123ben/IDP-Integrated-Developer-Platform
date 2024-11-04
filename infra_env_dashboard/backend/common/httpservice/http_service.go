package httpservice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// HttpService wraps the HTTP client and provides methods for making requests.
type HttpService struct {
	client *http.Client
}

// New creates a new instance of HttpService with a default HTTP client.
func New() *HttpService {
	return &HttpService{
		client: &http.Client{},
	}
}

// Get makes a GET request to the specified URL and decodes the response into the provided output.
func (s *HttpService) Get(url string, token string, output interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set Authorization header if a token is provided
	if token != "" {
		req.Header.Set("Authorization", "token "+token)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(output)
}
