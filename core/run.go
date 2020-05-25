package core

// RunConfig holds details about the loadgen execution
type RunConfig struct {
	TPS      int
	Duration int
}

// RunSummary holds details about the load test results
type RunSummary struct {
	RunID    string
	Requests RequestSummary
	Latency  map[string]int
}

// RequestSummary is a breakdown of the requests
type RequestSummary struct {
	ErrorCount   int
	SuccessCount int
	TotalCount   int
	ByStatusCode map[string]int
}

// Run starts a load test with the given run configuration
// and an object that conforms to the Hammer interface.
// Returns a unique id for the execution.
func Run(config RunConfig, h Hammer) RunSummary {

	// Generate uuid for execution
	id := uuid()

	stop := make(chan bool)
	done, responses := loadgen(config, h)
	s := aggregate(responses, stop)

	<-done       // wait until loadgen finishes
	stop <- true // tell aggregator to stop

	summary := <-s // read summary from aggregator
	summary.RunID = id
	printSummary(summary)

	return summary
}
