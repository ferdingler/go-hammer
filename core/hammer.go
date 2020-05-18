package core

import (
	"bytes"
	"net/http"
	"sync"
	"time"
)

// Hammer defines functions to be implemented by hammers
type Hammer interface {
	Hit() HammerResponse
}

// HammerResponse is information about a hammer response
type HammerResponse struct {
	Latency   int // milliseconds
	Status    int
	Timestamp time.Time
	Failed    bool
}

func useHammer(h Hammer, out chan HammerResponse, wg *sync.WaitGroup) {
	response := h.Hit()
	out <- response
	defer wg.Done()
}

// Built-in Hammers
//
// The following are definitions and implementations of the
// built-in hammers this library offers. Anyone can develop
// their own custom hammers but for quick usage, they can use
// the built-in ones provided below.

// HTTPHammer built-in for http requests
type HTTPHammer struct {
	client      *http.Client
	request     *http.Request
	Endpoint    string
	Method      string
	ContentType string
	Body        []byte
	Headers     map[string]string
}

// Hit method for HTTPHammer
func (h *HTTPHammer) Hit() HammerResponse {
	if h.client == nil {
		h.client = new(http.Client)
		h.client.Timeout = time.Second * 10
	}

	// Trigger HTTP request and time it
	start := time.Now()
	res, err := httpRequest(h)
	end := time.Now()
	diff := end.Sub(start)

	if err != nil {
		// non-2xx response doesn't cause an error,
		// so this error means something bad happened.
		return HammerResponse{
			Latency:   0,
			Status:    0,
			Timestamp: start.UTC(),
			Failed:    true,
		}
	}

	// close response body
	defer res.Body.Close()
	return HammerResponse{
		Latency:   int(diff.Milliseconds()),
		Status:    res.StatusCode,
		Timestamp: start.UTC(),
	}
}

func httpRequest(h *HTTPHammer) (*http.Response, error) {
	if h.request == nil {
		body := bytes.NewBuffer(h.Body)
		req, err := http.NewRequest(h.Method, h.Endpoint, body)
		if err != nil {
			panic("Invalid HTTP request")
		}

		if len(h.ContentType) > 0 {
			req.Header.Add("Content-Type", h.ContentType)
		}

		if len(h.Headers) > 0 {
			for key, value := range h.Headers {
				req.Header.Add(key, value)
			}
		}

		h.request = req
	}

	return h.client.Do(h.request)
}
