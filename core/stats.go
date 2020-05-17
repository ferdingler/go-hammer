package core

import "sort"

type runSummary struct {
	p99 int
	p95 int
	p90 int
	p50 int
}

func summarize(latencies []int) runSummary {
	sort.Ints(latencies)
	return runSummary{
		p99: percentile(latencies, 99),
		p95: percentile(latencies, 95),
		p90: percentile(latencies, 90),
		p50: percentile(latencies, 50),
	}
}

func percentile(values []int, p float32) int {
	if len(values) == 0 {
		return 0
	}

	rank := int((p / 100) * float32(len(values)+1))
	return values[rank-1]
}
