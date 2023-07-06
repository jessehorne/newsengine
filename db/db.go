package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var Ready bool

func InitDB() error {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	sslMode := os.Getenv("DB_SSL")

	connStr := fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=%v",
		username, password, dbHost, dbName, sslMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	DB = db

	return nil
}
