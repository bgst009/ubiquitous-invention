package main

import (
	"fmt"

	"github.com/bgst009/ubiquitous-invention/internal/app/monitor"
)

func main() {
	fmt.Println("Starting")

	// Create a new instance of the service
	monitor.TickM()
}
