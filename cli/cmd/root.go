package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ferdingler/go-hammer/core"
	"github.com/ferdingler/go-hammer/core/hammers"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var endpoint string
var tps int
var duration int
var method string
var payload string
var headers string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gohammer",
	Short:   "A lightweight load testing tool",
	Version: "v0.1.7",
	Long: `
go-hammer is a load testing tool written in Go.
This CLI serves as a quick way to run load tests 
by invoking commands without writing code.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		config := core.RunConfig{
			TPS:      tps,
			Duration: duration,
		}

		hammer := new(hammers.HTTPHammer)
		hammer.Endpoint = endpoint
		hammer.Method = method
		hammer.Body = []byte(payload)

		if len(headers) > 0 {
			var headersMap map[string]string
			json.Unmarshal([]byte(headers), &headersMap)
			hammer.Headers = headersMap
		}

		core.Run(config, hammer)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml")
	rootCmd.Flags().StringVarP(&endpoint, "endpoint", "", "", "Specify HTTP target endpoint")
	rootCmd.Flags().IntVarP(&tps, "tps", "", 1, "Number of requests per second")
	rootCmd.Flags().IntVarP(&duration, "duration", "", 60, "Desired duration in seconds")
	rootCmd.Flags().StringVarP(&method, "method", "", "GET", "HTTP method")
	rootCmd.Flags().StringVarP(&payload, "payload", "", "", "Payload body for the HTTP requests")
	rootCmd.Flags().StringVarP(&headers, "headers", "", "", "HTTP headers to include on each request")

	// Mark required flags
	rootCmd.MarkFlagRequired("endpoint")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
