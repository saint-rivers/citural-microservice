package database

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/saint-rivers/tinker/api/container"
	"github.com/saint-rivers/tinker/common"
)

func CreateService(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req request.DatabaseServiceRequest
		json.NewDecoder(r.Body).Decode(&req)

		row := InsertDatabaseRequest(db, &req)

		if row.Err() != nil {
			var errorResponse = map[string]string{
				"message":   "unable to create database service",
				"timestamp": time.Now().Local().Format("2006-01-02T15:04:05Z07:00"),
			}
			json.NewEncoder(w).Encode(errorResponse)
			log.Println(row.Err().Error())
			return
		}

		c := container.CreateDatabaseContainer(&req)
		container.StartDatabaseContainer(c.ID)

		json.NewEncoder(w).Encode(req)
	}
}
