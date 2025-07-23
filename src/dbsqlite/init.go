package dbsqlite

import (
	"database/sql"
	"fmt"

	"github.com/adrg/xdg"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() error {
	dbFile, err := FileDB()
	if err != nil {
		return fmt.Errorf("could not get database file path: %w", err)
	}

	conn, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return fmt.Errorf("could not open database: %w", err)
	}
	db = conn

	if err = createTables(); err != nil {
		return fmt.Errorf("could not create tables: %w", err)
	}

	return nil
}

func DB() *sql.DB {
	return db
}

func createTables() error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS contacts (id INTEGER PRIMARY KEY, name TEXT, whatsapp_number TEXT)")
	if err != nil {
		return err
	}
	return nil
}

func FileDB() (string, error) {
	configFilePath, err := xdg.ConfigFile("wa-blazterr/database.db")
	if err != nil {
		return "", fmt.Errorf("could not resolve path for database file: %w", err)
	}
	return configFilePath, nil
}

// Close closes the database connection.
func Close() {
	if db != nil {
		db.Close()
	}
}
