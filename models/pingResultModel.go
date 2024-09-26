package models

import (
	"serverStateService/entities"
	"serverStateService/utils"
)

type PingResultModel struct {
	db *utils.DatabaseHandler
}

func (m *PingResultModel) GetResults() (*[]entities.PingResult, error) {
	err := m.db.Connect()
	if err != nil {
		return nil, err
	}

	rows, err := m.db.DB.Query("SELECT id, idServer, scannedAt FROM pingResults")
	if err != nil {
		return nil, err
	}

	var pings []entities.PingResult

	for rows.Next() {
		var ping entities.PingResult

		err = rows.Scan(&ping)
		if err != nil {
			return nil, err
		}

		pings = append(pings, ping)
	}

	err = m.db.Disconnect()
	if err != nil {
		return nil, err
	}

	return &pings, nil
}
