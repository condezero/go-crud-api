package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	// registering db driver
	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	uri := os.Getenv("DATABASE_CONNECTIONSTRING")
	fmt.Printf(uri)
	return sql.Open("postgres", uri)
}

// MakeMigration creates all the tables in the database
func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./migrations/migration.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
