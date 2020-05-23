package hammers

import (
	"bytes"
	"net/http"
	"time"
)

// HammerResponse is information about a hammer response
type HammerResponse struct {
	Latency   int // milliseconds
	Status    int
	Timestamp time.Time
	Failed    bool
}

// HTTPHammer built-in for http requests
type HTTPHammer struct {
	client      *http.Client
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

	return h.client.Do(req)
}
