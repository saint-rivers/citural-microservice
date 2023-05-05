package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/saint-rivers/tinker/api/database"
	"github.com/saint-rivers/tinker/api/health"
)

func RegisterRoutes(db *sql.DB, router *mux.Router) {
	health.RegisterHealthcheckRoutes(db, router)
	database.RegisterDatabaseServiceRoutes(db, router)
}