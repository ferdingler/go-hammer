package core

// RunConfig holds details about the loadgen execution
type RunConfig struct {
	TPS      int
	Duration int
}

// Run starts a load test with the given run configuration
// and an object that conforms to the Hammer interface.
// Returns a unique id for the execution.
func Run(config RunConfig, h Hammer) string {

	// Generate uuid for execution
	id := uuid()

	stop := make(chan bool)
	done, responses := loadgen(config, h)
	s := aggregate(responses, stop)

	<-done         // wait until loadgen finishes
	stop <- true   // tell aggregator to stop
	summary := <-s // read summary from aggregator

	outSummary(summary)
	return id
}
