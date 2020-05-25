# go-hammer

<img src="docs/logo.svg" align="right" height="140">

![Build](https://github.com/ferdingler/go-hammer/workflows/Build/badge.svg)

A load generator written in Go. 
<br><br>

## Usage

### Command Line Tool

The easiest way to get started is by using the CLI.

```bash
cli --endpoint https://www.google.com --duration 60 --tps 1
```

To install the CLI, run `go install` within the _cli_ folder.

### Core Library

Alternatively, you can use the core library directly to write your own load tests in Go. 

```go
import (
	"github.com/ferdingler/go-hammer/core"
	"github.com/ferdingler/go-hammer/hammers"
)

func main() {
	config := core.RunConfig{
		TPS:      10,
		Duration: 60,
	}

	hammer := new(hammers.HTTPHammer)
	hammer.Endpoint = "https://www.google.com"
	hammer.Method = "GET"
	hammer.Headers = map[string]string{
		"content-type": "application/json",
	}

	core.Run(config, hammer)
}
```

## What is a Hammer?

The name go-hammer comes from the analogy of comparing load testing a service with the activity of hitting a nail hard with a hammer many times. The nail would be your endpoint and the hammer is the tool used to hit it. 

This concept provides extensibility to the library by decoupling the load generation from the logic that creates each request (the hammer). In other words, you can swap the hammer that you use to _hit_ your endpoint with any hammer that conforms to the `Hammer` interface. 

```go
type Hammer interface {
	Hit() HammerResponse
}

type HammerResponse struct {
	Latency   int // milliseconds
	Status    int // status code
	Timestamp time.Time
	Failed    bool
	Body      []byte
}
```

The `hammers` package contains built-in hammers that you can leverage to get started quickly, like the HTTPHammer. But you can write your own custom _hammers_ that have any logic you want to run on every request.