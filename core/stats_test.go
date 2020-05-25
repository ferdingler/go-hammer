package core

import (
	"testing"
	"time"
)

func TestP50(t *testing.T) {
	values := getValues()
	p := percentile(values, 50)
	if p != 7 {
		t.Errorf("p50 should be 7, got %d", p)
	}
}

func TestP90(t *testing.T) {
	values := getValues()
	p := percentile(values, 90)
	if p != 15 {
		t.Errorf("p90 should be 15, got %d", p)
	}
}

func TestP99(t *testing.T) {
	values := getValues()
	p := percentile(values, 99)
	if p != 15 {
		t.Errorf("p90 should be 15, got %d", p)
	}
}

func TestP100(t *testing.T) {
	values := getValues()
	p := percentile(values, 100)
	if p != 23 {
		t.Errorf("p100 should be 23, got %d", p)
	}
}

func TestSummarize(t *testing.T) {
	var results []HammerResponse

	results = append(results, HammerResponse{
		Latency:   10,
		Status:    201,
		Timestamp: time.Now(),
		Failed:    false,
	})

	results = append(results, HammerResponse{
		Latency:   5,
		Status:    201,
		Timestamp: time.Now(),
		Failed:    false,
	})

	results = append(results, HammerResponse{
		Latency:   100,
		Status:    403,
		Timestamp: time.Now(),
		Failed:    false,
	})

	results = append(results, HammerResponse{
		Latency:   12,
		Status:    403,
		Timestamp: time.Now(),
		Failed:    false,
	})

	results = append(results, HammerResponse{
		Latency:   0,
		Status:    0,
		Timestamp: time.Now(),
		Failed:    true,
	})

	summary := summarize(results)
	requests := summary.Requests
	if requests.TotalCount != 5 {
		t.Errorf("Expected summary to have 5 total requests, got %d", requests.TotalCount)
	}

	if requests.ErrorCount != 1 {
		t.Errorf("Expected summary to have 1 error requests, got %d", requests.ErrorCount)
	}

	if requests.SuccessCount != 4 {
		t.Errorf("Expected summary to have 4 success requests, got %d", requests.SuccessCount)
	}

	count, exists := requests.ByStatusCode["201"]
	if !exists {
		t.Error("Expected summary to have status code 201")
	}

	if count != 2 {
		t.Errorf("Expected summary to have 2 status code 201, got %d", count)
	}

	count, exists = requests.ByStatusCode["403"]
	if !exists {
		t.Error("Expected summary to have status code 403")
	}

	if count != 2 {
		t.Errorf("Expected summary to have 2 status code 403, got %d", count)
	}

	// I don't care about testing the percentile values
	expectedPercentiles := map[string]float64{
		"p100":  0,
		"p99.9": 0,
		"p99":   0,
		"p95":   0,
		"p90":   0,
		"p50":   0,
	}

	for percentile := range expectedPercentiles {
		_, exists := summary.Latency[percentile]
		if !exists {
			t.Errorf("Expected summary to include percentile %s", percentile)
		}
	}
}

func getValues() []int {
	return []int{
		1,
		1,
		1,
		3,
		3,
		5,
		5,
		7,
		8,
		9,
		10,
		11,
		11,
		13,
		15,
		15,
		23,
	}
}
