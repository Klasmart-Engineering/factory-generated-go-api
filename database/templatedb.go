package db

import (
	"database/sql"
	"fmt"
	"go-template/models"
)

func (db Database) Healthcheck() (bool, error) {
	var up int
	if err := db.Conn.QueryRow(`
			SELECT 1 AS up 
			FROM go_template_table`).Scan(&up); err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("unable to get connection to the database: %s", err)
		}
	}

	return up == 1, nil
}

func (db Database) ExampleGet() (*models.ExampleStructList, error) {
	list := &models.ExampleStructList{}
	rows, err := db.Conn.Query("SELECT id, name FROM go_template_table ORDER BY id DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var example models.ExampleStruct
		err := rows.Scan(&example.ExampleId, &example.ExampleName)
		if err != nil {
			return list, err
		}
		list.ExampleStruct = append(list.ExampleStruct, example)
	}
	return list, nil
}
