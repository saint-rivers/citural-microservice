package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/saint-rivers/tinker/api/v1/healthcheck"
	"github.com/saint-rivers/tinker/api/v1/services"
)

func EnableHealthcheckRoutes(db *sql.DB, router *mux.Router) {
	router.HandleFunc("/healthcheck", healthcheck.HealthCheck()).Methods("GET")
}

func EnableServiceRoutes(db *sql.DB, router *mux.Router) {
	router.HandleFunc("/api/v1/services", services.CreateService(db)).Methods("POST")
}
