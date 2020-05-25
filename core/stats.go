package core

import (
	"sort"
	"strconv"
)

func summarize(results []HammerResponse) RunSummary {

	errors := 0
	var latencies []int
	countByStatus := make(map[string]int)
	for _, res := range results {
		if res.Failed {
			errors++
		} else {
			latencies = append(latencies, res.Latency)
			code := strconv.Itoa(res.Status)
			countByStatus[code]++
		}
	}

	sort.Ints(latencies)
	latencyMap := map[string]int{
		"p100":  percentile(latencies, 100),
		"p99.9": percentile(latencies, 99.9),
		"p99":   percentile(latencies, 99),
		"p95":   percentile(latencies, 95),
		"p90":   percentile(latencies, 90),
		"p50":   percentile(latencies, 50),
	}

	return RunSummary{
		Latency: latencyMap,
		Requests: RequestSummary{
			ErrorCount:   errors,
			SuccessCount: len(results) - errors,
			TotalCount:   len(results),
			ByStatusCode: countByStatus,
		},
	}
}

func percentile(values []int, p float32) int {
	if len(values) == 0 {
		return 0
	}

	rank := int((p / 100) * float32(len(values)))
	return values[rank-1]
}
