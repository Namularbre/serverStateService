package utils

import (
	"database/sql"
	"os"
)

type DatabaseHandler struct {
	DB *sql.DB
}

func (d *DatabaseHandler) Connect() error {
	db, err := sql.Open("mysql", os.Getenv("CONN_STR"))
	d.DB = db
	return err
}

func (d *DatabaseHandler) Disconnect() error {
	if d.DB != nil {
		return d.DB.Close()
	}
	return nil
}
