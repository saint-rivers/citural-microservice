package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/saint-rivers/tinker/api/database"
	"github.com/saint-rivers/tinker/api/health"
)

func EnableHealthcheckRoutes(db *sql.DB, router *mux.Router) {
	router.HandleFunc("/healthcheck", health.HealthCheck()).Methods("GET")
}

func EnableServiceRoutes(db *sql.DB, router *mux.Router) {
	router.HandleFunc("/api/v1/services", database.CreateService(db)).Methods("POST")
}
