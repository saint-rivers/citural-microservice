package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type DatabaseServiceRequest struct {
	ContainerName   string `json:"containerName"`
	DatabaseVendor  string `json:"vendor"`
	DefaultDatabase string `json:"defaultDatabase"`
	UserName        string `json:"username"`
	Password        string `json:"password"`
	Port            string `json:"port"`
}

type DatabaseService struct {
	Id              int64  `json:"id"`
	ContainerName   string `json:"containerName"`
	DatabaseVendor  string `json:"vendor"`
	DefaultDatabase string `json:"defaultDatabase"`
	UserName        string `json:"username"`
	Password        string `json:"password"`
	Port            string `json:"port"`
}

func CreateService(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dbRequest DatabaseServiceRequest
		json.NewDecoder(r.Body).Decode(&dbRequest)
		fmt.Println(dbRequest)

		err := db.QueryRow(
			`INSERT INTO app_service (container_name, database_vendor, default_database, db_username, db_password)
			VALUES ($1, $2, $3, $4, $5);
			`,
			dbRequest.ContainerName, dbRequest.DatabaseVendor, dbRequest.DefaultDatabase, dbRequest.UserName, dbRequest.Password,
		)
		if err.Err() != nil {
			var errorResponse = map[string]string{
				"message":   "unable to create database service",
				"timestamp": time.Now().Local().Format("2006-01-02T15:04:05Z07:00"),
			}
			json.NewEncoder(w).Encode(errorResponse)
			log.Println(err.Err().Error())
			return
		}

		fmt.Print(err)
		json.NewEncoder(w).Encode(dbRequest)
	}
}
