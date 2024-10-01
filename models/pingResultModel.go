package models

import (
	"database/sql"
	"os"
	"serverStateService/entities"
)

func InsertResult(result *entities.PingResult) error {
	db, err := sql.Open("mysql", os.Getenv("CONN_STR"))
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO pingResults (idServer, scannedAt, info) VALUES (?, ?, ?);", result.IdServer, result.ScannedAt, result.Info)
	return err
}
