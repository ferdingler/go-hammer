package core

import (
	"testing"
	"time"
)

func TestAggregator(t *testing.T) {
	// Aggregator listens for responses in a channel
	responses := make(chan HammerResponse)
	// And listens for a stop signal
	stop := make(chan bool)

	// Run aggregator, it gives back a channel
	// where the summary will be returned.
	// Since it runs in a goroutine, need a timeout
	// to fail the test if it exceeds the time.
	timeout := time.After(5 * time.Second)
	results := aggregate(responses, stop)

	// Generate and send some responses to aggregator
	responses <- HammerResponse{
		Latency:   1,
		Status:    200,
		Timestamp: time.Now(),
		Failed:    false,
	}

	responses <- HammerResponse{
		Latency:   2,
		Status:    200,
		Timestamp: time.Now(),
		Failed:    false,
	}

	// Stop aggregator
	stop <- true

	// Collect summary
	select {
	case summary := <-results:
		if summary.p99 <= 0 {
			t.Errorf("Expected p99 to be greater than zero, got %d", summary.p99)
		}
	case <-timeout:
		t.Error("Test timeout exceeded")
	}
}
