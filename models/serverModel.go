package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"serverStateService/entities"
)

func GetServers() ([]entities.Server, error) {
	db, err := sql.Open("mysql", os.Getenv("CONN_STR"))
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, ip, name FROM servers")
	if err != nil {
		return nil, err
	}

	var servers []entities.Server

	for rows.Next() {
		var id int
		var ip string
		var hostname string

		err = rows.Scan(&id, &ip, &hostname)
		if err != nil {
			return nil, err
		}

		server := entities.Server{
			Id:   id,
			Ip:   ip,
			Name: hostname,
		}

		servers = append(servers, server)
	}

	return servers, nil
}
