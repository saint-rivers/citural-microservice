package database

import (
	"database/sql"

	"github.com/gorilla/mux"
	sec "github.com/saint-rivers/tinker/api/auth"
)

func RegisterDatabaseServiceRoutes(db *sql.DB, router *mux.Router) {
	router.Handle("/api/v1/services", ListContainers()).Methods("GET")
	router.Handle("/api/v1/services", sec.Protect(CreateService(db))).Methods("POST")
	router.Handle("/api/v1/services", sec.Protect(ManageService())).Methods("PUT")
	router.Handle("/api/v1/services", sec.Protect(RemoveContainer())).Methods("DELETE")
}
