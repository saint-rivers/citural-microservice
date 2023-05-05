package health

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterHealthcheckRoutes(db *sql.DB, router *mux.Router) {
	router.HandleFunc("/healthcheck", HealthCheck()).Methods("GET")
}
