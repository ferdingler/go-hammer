package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

// Writing a new http hammer to avoid having a dependency
// on the hammers package.
type httpHammer struct {
	client      *http.Client
	Endpoint    string
	ContentType string
	Method      string
	Body        []byte
}

func (h *httpHammer) Hit() HammerResponse {
	if h.client == nil {
		h.client = new(http.Client)
	}

	body := bytes.NewBuffer(h.Body)
	req, _ := http.NewRequest(h.Method, h.Endpoint, body)
	res, err := h.client.Do(req)
	if err != nil {
		panic("Error with request")
	}

	defer res.Body.Close()
	return HammerResponse{
		Latency:   1,
		Status:    res.StatusCode,
		Timestamp: time.Now(),
	}
}

// This is essentially a functional test, not a unit test.
// It stands up a webserver and starts a loadgen execution
// against it. It verifies that the expected requests are
// actually received by the server.
func TestLoadGen(t *testing.T) {

	// Execution values
	config := RunConfig{}
	config.TPS = 1
	config.Duration = 10

	body, _ := json.Marshal(map[string]string{
		"message": "Hello world",
	})

	// Use the built-in http hammer
	hammer := new(httpHammer)
	hammer.Endpoint = "http://127.0.0.1:3000/foo"
	hammer.Method = "PUT"
	hammer.ContentType = "application/json"
	hammer.Body = body

	// Expected values
	requests := 0
	totalExpected := config.TPS * config.Duration

	// Prepare web server
	done := make(chan bool)
	timeout := time.After(20 * time.Second)
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		// Incrementing this global vairable is a bad idea because
		// there may be race conditions if tps > 1. But for now is fine
		// as I only have 1 tps.
		requests++

		var msg map[string]string
		json.NewDecoder(r.Body).Decode(&msg)

		// Assertions
		if r.Method != "PUT" {
			t.Error("Expected http method to be PUT")
		}

		if msg["message"] != "Hello world" {
			t.Error("Expected body to have a hello world message")
		}

		if requests == totalExpected {
			// We are done
			fmt.Printf("Total requests received %d\n", requests)
			done <- true
		}
	})

	// Start listening for requests
	go http.ListenAndServe(":3000", nil)
	time.Sleep(2 * time.Second)

	// Start load test
	go Run(config, hammer)

	// This select is useful to know if we have finished
	// receiving all requests or if we timeout.
	select {
	case <-done:
		fmt.Println("Test completed")
	case <-timeout:
		t.Error("Test timeout")
		fmt.Printf("Requests received %d\n", requests)
	}
}
