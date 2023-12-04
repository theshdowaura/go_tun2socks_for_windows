package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDistinctFilenames() ([]string, error) {
	db, err := sql.Open("sqlite3", "database/rules.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT DISTINCT filename FROM rules")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var filenames []string
	for rows.Next() {
		var filename string
		if err := rows.Scan(&filename); err != nil {
			return nil, err
		}
		filenames = append(filenames, filename)
	}
	return filenames, nil
}
