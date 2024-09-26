package models

import (
	"serverStateService/entities"
	"serverStateService/utils"

	_ "github.com/go-sql-driver/mysql"
)

type ServerModel struct {
	db *utils.DatabaseHandler
}

func (m *ServerModel) GetServers() ([]entities.Server, error) {
	err := m.db.Connect()
	if err != nil {
		return nil, err
	}

	rows, err := m.db.DB.Query("SELECT id, ip, name FROM servers")
	if err != nil {
		return nil, err
	}

	var servers []entities.Server

	for rows.Next() {
		var server entities.Server

		err = rows.Scan(&server)
		if err != nil {
			return nil, err
		}

		servers = append(servers, server)
	}

	err = m.db.Disconnect()
	if err != nil {
		return nil, err
	}

	return servers, nil
}
