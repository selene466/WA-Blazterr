package dbinit

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/adrg/xdg"
	_ "github.com/mattn/go-sqlite3"
)

type ConfigDB struct {
	dbPath string
}

func Init() {
	dbFile, err := FileDB()
	if err != nil {
		fmt.Printf("Could not initialize the database file: %v\n", err)
		return
	}
	db, err := sql.Open("sqlite3", dbFile.dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS contacts (id INTEGER PRIMARY KEY, name TEXT, whatsapp_number TEXT)")
	if err != nil {
		log.Fatal(err)
	}
}

func FileDB() (*ConfigDB, error) {
	configFilePath, err := xdg.ConfigFile("wa-iris/database.db")
	if err != nil {
		return nil, fmt.Errorf("COULD NOT RESOLVE PATH FOR DATABASE FILE: %w", err)
	}

	return &ConfigDB{
		dbPath: configFilePath,
	}, nil
}

func (c *ConfigDB) FileDBPath() string {
	return c.dbPath
}
