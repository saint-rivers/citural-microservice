package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/saint-rivers/tinker/env"
	"log"
)

func GetDBConnection() *sql.DB {
	db, err := sql.Open("postgres", env.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	return db
}
