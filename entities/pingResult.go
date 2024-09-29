package entities

import "time"

type PingResult struct {
	Id        int        `json:"id"`
	IdServer  int        `json:"idServer"`
	ScannedAt *time.Time `json:"scannedAt"`
	Respond   bool
}
