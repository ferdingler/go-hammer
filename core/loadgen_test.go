package core

import (
	"testing"
	"time"
)

type dumbHammer struct {
	timesInvoked int
}

func (h *dumbHammer) Hit() HammerResponse {
	h.timesInvoked++
	return HammerResponse{
		Latency: 1,
		Status:  200,
	}
}

func TestLoadgen(t *testing.T) {
	// Loadgen receives a Run configuration
	config := RunConfig{
		Duration: 3, // Five seconds
		TPS:      5,
	}

	// and receives an instance of a hammer
	hammer := new(dumbHammer)

	// Run the load generator, it gives back 2 channels
	done, responses := loadgen(config, hammer)

	// Since it is a goroutine, we need a timeout
	// to fail the test if it exceeds the time.
	timeout := time.After(5 * time.Second)

	// Collect summary
	for {
		select {
		case <-done:
			actual := hammer.timesInvoked
			expected := config.Duration * config.TPS
			if actual != expected {
				t.Errorf("Expected invocations to be %d, got %d", expected, actual)
			}
			return
		case <-responses:
			// The only reason to read the responses channel is to
			// unblock the hammers, otherwise they wait forever
			// until their message is read.
			continue
		case <-timeout:
			t.Error("Test timeout exceeded")
		}
	}
}
