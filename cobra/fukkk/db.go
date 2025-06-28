package fukkk

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Dbinit() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "fukkk/task.db")
	if err != nil {
		return nil, err
	}
	return db, err
}

func Logs(db *sql.DB, input string, output string, date string) error {

	zapr, err := db.Prepare("INSERT INTO logs(input, output, date) VALUES(?, ?,?)")

	if err != nil {
		return err
	}

	defer zapr.Close()
	_, err = zapr.Exec(input, output, date)
	return err

}

func CreateTable(db *sql.DB) error {

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			input TEXT,
			output TEXT,
			date TEXT
		)
		`)
	return err
}

type GETALLLOGS struct {
	Date string `json:"date"`
	Task string `json:"task"`
}

func GetAllLogs(db *sql.DB) ([]GETALLLOGS, error) {
	rows, err := db.Query("SELECT date, input  FROM logs ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []GETALLLOGS
	for rows.Next() {
		var input GETALLLOGS
		if err := rows.Scan(&input.Date, &input.Task); err != nil {
			return nil, err
		}
		logs = append(logs, input)
	}

	return logs, nil
}
