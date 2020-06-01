/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ferdingler/go-hammer/core"
	"github.com/ferdingler/go-hammer/hammers"
	"github.com/spf13/cobra"
)

var s3Bucket string
var s3Key string
var path string

// Scenario definition
type Scenario struct {
	Duration     int             `json:"duration"`
	TPS          int             `json:"tps"`
	Hammer       string          `json:"hammer"`
	HammerConfig json.RawMessage `json:"hammerConfig"`
}

// HttpConfig
type HttpConfig struct {
	Endpoint string            `json:"endpoint"`
	Method   string            `json:"method"`
	Payload  string            `json:"payload"`
	Headers  map[string]string `json:"headers"`
}

// scenarioCmd represents the scenario command
var scenarioCmd = &cobra.Command{
	Use:   "scenario",
	Short: "Run a load test scenario defined in JSON",
	Run: func(cmd *cobra.Command, args []string) {

		// Read scenario from local path
		data, err := ioutil.ReadFile(path)
		check(err)

		var scenario Scenario
		err = json.Unmarshal(data, &scenario)
		check(err)

		// Parse the Hammer Config
		var httpConfig HttpConfig
		err = json.Unmarshal(scenario.HammerConfig, &httpConfig)
		check(err)

		// Start running HTTP hammer
		config := core.RunConfig{
			TPS:      scenario.TPS,
			Duration: scenario.Duration,
		}

		hammer := new(hammers.HTTPHammer)
		hammer.Endpoint = httpConfig.Endpoint
		hammer.Method = httpConfig.Method
		hammer.Body = []byte(httpConfig.Payload)
		hammer.Headers = httpConfig.Headers

		fmt.Println(httpConfig)
		core.Run(config, hammer)
	},
}

func init() {
	runCmd.AddCommand(scenarioCmd)

	scenarioCmd.Flags().StringVar(&s3Bucket, "s3-bucket", "", "S3 bucket where scenario file is stored")
	scenarioCmd.Flags().StringVar(&s3Key, "s3-key", "", "S3 object key where scenario file is stored")
	scenarioCmd.Flags().StringVar(&path, "path", "", "Local path to scenario file")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
