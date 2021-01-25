package database

import (
	"database/sql"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	uri := os.Getenv("POSTGRESQL_URI")
	return sql.Open("postgres", uri)
}
func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("migrations/migration.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
