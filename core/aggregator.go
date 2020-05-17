package core

import (
	"fmt"
)

func aggregate(responses chan HammerResponse, stop chan bool) chan runSummary {
	summary := make(chan runSummary)
	go func() {
		// Holds values in-memory
		var latencies []int
		for {
			select {
			// Continue reading from responses until signaled
			// to stop by the channel.
			case response := <-responses:
				latencies = append(latencies, response.Latency)
				outResponse(response)
			// Signal to stop by the main routine,
			// compute summary and report it back
			case <-stop:
				fmt.Println("Aggregator finished, summarizing")
				summary <- summarize(latencies)
			}
		}
	}()
	return summary
}
