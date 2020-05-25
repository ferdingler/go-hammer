package core

import (
	"encoding/json"
	"fmt"
)

func printResponse(response HammerResponse) {
	fmt.Printf("%s,%d,%d\n",
		response.Timestamp,
		response.Status,
		response.Latency,
	)
}

func printSummary(summary RunSummary) {
	jsonData, _ := json.MarshalIndent(summary, "", "  ")
	fmt.Println(string(jsonData))
}
