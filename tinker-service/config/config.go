package config

import (
	"database/sql"
	"log"

	"github.com/docker/docker/client"
	_ "github.com/lib/pq"
	"github.com/saint-rivers/tinker/env"
)

func GetDBConnection() *sql.DB {
	db, err := sql.Open("postgres", env.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	return db
}

func GetDockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	return cli
}