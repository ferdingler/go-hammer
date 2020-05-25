package core

import (
	"fmt"
)

func aggregate(responses chan HammerResponse, stop chan bool) chan RunSummary {
	summary := make(chan RunSummary)
	go func() {
		// Hold values in-memory
		var results []HammerResponse
		for {
			select {
			// Continue reading from responses until signaled
			// to stop by the channel.
			case response := <-responses:
				results = append(results, response)
				printResponse(response)
			// Signal to stop by the main routine,
			// compute summary and report it back
			case <-stop:
				fmt.Println("Aggregator finished, summarizing")
				summary <- summarize(results)
			}
		}
	}()
	return summary
}
