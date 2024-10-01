package entities

import "time"

type PingResult struct {
	Id        int
	IdServer  int
	ScannedAt *time.Time
	Info      string
}
