module github.com/ferdingler/go-hammer/cli

go 1.13

replace github.com/ferdingler/go-hammer/core => /Users/fdingler/go/src/github.com/ferdingler/go-hammer/core

replace github.com/ferdingler/go-hammer/hammers => /Users/fdingler/go/src/github.com/ferdingler/go-hammer/hammers

require (
	github.com/ferdingler/go-hammer/core v0.1.7
	github.com/ferdingler/go-hammer/hammers v0.1.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
)
