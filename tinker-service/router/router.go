package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	sec "github.com/saint-rivers/tinker/api/auth"
	"github.com/saint-rivers/tinker/api/database"
	"github.com/saint-rivers/tinker/api/health"
)

func EnableHealthcheckRoutes(db *sql.DB, router *mux.Router) {
	router.HandleFunc("/healthcheck", health.HealthCheck()).Methods("GET")
}

func EnableServiceRoutes(db *sql.DB, router *mux.Router) {
	// router.HandleFunc("/api/v1/services", database.CreateService(db)).Methods("POST")
	router.Handle("/api/v1/services", sec.Protect(database.CreateService(db))).Methods("POST")
}
