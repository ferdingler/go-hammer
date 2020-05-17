package core

import (
	"fmt"
)

func outResponse(response HammerResponse) {
	fmt.Printf("%s,%d,%d\n",
		response.Timestamp,
		response.Status,
		response.Latency,
	)
}

func outSummary(summary runSummary) {
	fmt.Println("--")
	fmt.Println("Summary")
	fmt.Println("--")
	fmt.Printf("p99: %d\n", summary.p99)
	fmt.Printf("p95: %d\n", summary.p95)
	fmt.Printf("p90: %d\n", summary.p90)
	fmt.Printf("p50: %d\n", summary.p50)
}
