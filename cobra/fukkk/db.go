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

func Logs(db *sql.DB, input string, output string) error {

	zapr, err := db.Prepare("INSERT INTO logs(input, output) VALUES(?, ?)")

	if err != nil {
		return err
	}

	defer zapr.Close()
	_, err = zapr.Exec(input, output)
	return err

}

func CreateTable(db *sql.DB) error {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		input TEXT,
		output TEXT
	)
`)
	return err

}
