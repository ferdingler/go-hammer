package core

import (
	"sync"
	"testing"
	"time"
)

type bananaHammer struct{}

func (h *bananaHammer) Hit() HammerResponse {
	return HammerResponse{
		Latency: 100,
		Status:  500,
		Body:    []byte("Hello World"),
	}
}

func TestUseHammer(t *testing.T) {

	response := make(chan HammerResponse)
	hammer := new(bananaHammer)

	var wg sync.WaitGroup
	wg.Add(1)

	timeout := time.After(5 * time.Second)
	go useHammer(hammer, response, &wg)

	select {
	case r := <-response:
		if r.Latency != 100 {
			t.Errorf("Expected response latency to be 100, got %d", r.Latency)
		}

		if r.Status != 500 {
			t.Errorf("Expected response status to be 500, got %d", r.Status)
		}

		if string(r.Body) != "Hello World" {
			t.Errorf("Expected response body to be Hello World, got %d", r.Body)
		}
	case <-timeout:
		t.Error("Test timeout exceeded")
	}
}
