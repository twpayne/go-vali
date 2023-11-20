// Package vali provides a client interface to CIVL's Open Validation Server.
// See http://vali.fai-civl.org/webservice.html.
package vali

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

var defaultEndpoint = "http://vali.fai-civl.org/api/vali/json"

// A Status indicates the validity of an IGC file.
type Status int

const (
	StatusUnknown Status = iota // Unknown indicates that the validity of the IGC file is unknown.
	StatusValid                 // Valid indicates that the IGC file is valid.
	StatusInvalid               // Invalid indicates that the IGC file is invalid.
)

func (s Status) String() string {
	switch s {
	case StatusValid:
		return "Valid"
	case StatusInvalid:
		return "Invalid"
	case StatusUnknown:
		return "Unknown"
	default:
		return "Invalid status"
	}
}

// A Response represents a response from the server.
type Response struct {
	Result string `json:"result"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
	IGC    string `json:"igc"`
	Ref    string `json:"ref"`
	Server string `json:"server"`
}

func (r Response) Passed() bool {
	return r.Result == "PASSED"
}

// A ServerError represents a server error.
type ServerError struct {
	HTTPStatusCode int
	HTTPStatus     string
}

func (se *ServerError) Error() string {
	return fmt.Sprintf("vali: %d %s", se.HTTPStatusCode, se.HTTPStatus)
}

// An Option is an option for configuring a Service.
type Option func(*Service)

// Client sets the http.Client.
func Client(client *http.Client) Option {
	return func(s *Service) {
		s.client = client
	}
}

// Endpoint sets the HTTP endpoint.
func Endpoint(endpoint string) Option {
	return func(s *Service) {
		s.endpoint = endpoint
	}
}

// A Service represents a validator service.
type Service struct {
	client   *http.Client
	endpoint string
}

// New returns a new Service.
func New(options ...Option) *Service {
	s := &Service{
		client:   &http.Client{},
		endpoint: defaultEndpoint,
	}
	for _, o := range options {
		o(s)
	}
	return s
}

// ValidateIGC validates igcFile.
func (s *Service) ValidateIGC(ctx context.Context, filename string, igcFile io.Reader) (Status, *Response, error) {
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	fw, err := w.CreateFormFile("igcfile", filename)
	if err != nil {
		return StatusUnknown, nil, err
	}
	if _, err = io.Copy(fw, igcFile); err != nil {
		return StatusUnknown, nil, err
	}
	if err := w.Close(); err != nil {
		return StatusUnknown, nil, err
	}
	req, err := http.NewRequest(http.MethodPost, s.endpoint, b)
	if err != nil {
		return StatusUnknown, nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := s.client.Do(req)
	if err != nil {
		return StatusUnknown, nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return StatusUnknown, nil, err
	}
	if err := resp.Body.Close(); err != nil {
		return StatusUnknown, nil, err
	}
	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		return StatusUnknown, nil, &ServerError{
			HTTPStatusCode: resp.StatusCode,
			HTTPStatus:     resp.Status,
		}
	}
	var r Response
	if err := json.Unmarshal(body, &r); err != nil {
		return StatusUnknown, nil, err
	}
	if r.Passed() {
		return StatusValid, &r, nil
	}
	return StatusInvalid, &r, nil
}
