package database

import (
	"database/sql"
	"log"
	"sync"
)

var (
	data *Data
	once sync.Once
)

// Data Manage the connection
type Data struct {
	DB *sql.DB
}

func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}
	err = MakeMigration(db)
	if err != nil {
		log.Panic(err)
	}
	data = &Data{
		DB: db,
	}
}

// New return a new instance
func New() *Data {
	once.Do(initDB)

	return data
}

// Close closes the resources
func Close() error {
	if data == nil {
		return nil
	}

	return data.DB.Close()
}
