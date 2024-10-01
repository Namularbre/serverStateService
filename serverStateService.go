package main

import (
	"serverStateService/monitoring"
	"time"
)

func main() {
	for {
		monitoring.RunPings()
		time.Sleep(12 * time.Hour)
	}
}
