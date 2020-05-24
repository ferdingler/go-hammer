package core

import (
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
	Body      []byte
}

func useHammer(h Hammer, out chan HammerResponse, wg *sync.WaitGroup) {
	response := h.Hit()
	out <- response
	defer wg.Done()
}
