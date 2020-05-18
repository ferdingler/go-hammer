package cmd

import (
	"encoding/json"

	gohammer "github.com/ferdingler/go-hammer/core"
	"github.com/spf13/cobra"
)

var endpoint string
var tps int
var duration int
var method string
var payload string
var headers string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts a load test execution",
	Long: `Starts a load test execution against a 
	target endpoint with the desired TPS and Duration.`,
	Run: func(cmd *cobra.Command, args []string) {

		config := gohammer.RunConfig{
			TPS:      tps,
			Duration: duration,
		}

		hammer := new(gohammer.HTTPHammer)
		hammer.Endpoint = endpoint
		hammer.Method = method
		hammer.Body = []byte(payload)

		if len(headers) > 0 {
			var headersMap map[string]string
			json.Unmarshal([]byte(headers), &headersMap)
			hammer.Headers = headersMap
		}

		gohammer.Run(config, hammer)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Flags and configuration settings.
	runCmd.Flags().StringVarP(&endpoint, "endpoint", "", "", "Specify HTTP target endpoint")
	runCmd.Flags().IntVarP(&tps, "tps", "", 1, "Number of requests per second")
	runCmd.Flags().IntVarP(&duration, "duration", "", 60, "Desired duration in seconds")
	runCmd.Flags().StringVarP(&method, "method", "", "GET", "HTTP method")
	runCmd.Flags().StringVarP(&payload, "payload", "", "", "Payload body for the HTTP requests")
	runCmd.Flags().StringVarP(&headers, "headers", "", "", "HTTP headers to include on each request")

	// Mark required flags
	runCmd.MarkFlagRequired("endpoint")
}
