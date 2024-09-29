package models

import (
	"serverStateService/entities"
	"serverStateService/utils"
)

type PingResultModel struct {
	db *utils.DatabaseHandler
}

func (m *PingResultModel) InsertResult(result *entities.PingResult) error {
	err := m.db.Connect()
	if err != nil {
		return err
	}

	_, err = m.db.DB.Exec("INSERT INTO pingResults (idServer, respond) VALUES (?, ?);", result.IdServer, result.Respond)
	if err != nil {
		return err
	}

	return m.db.Disconnect()
}
