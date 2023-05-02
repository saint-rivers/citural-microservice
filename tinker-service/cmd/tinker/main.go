package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/saint-rivers/tinker/config"
	"github.com/saint-rivers/tinker/env"
	"github.com/saint-rivers/tinker/router"
)

func configureCors(router *mux.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{env.ALLOWED_ORIGINS},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
	})
	handler := c.Handler(router)
	return handler
}

func startServer(handler http.Handler) {
	fmt.Printf("listening on %s", env.LISTEN_ADDRESS)
	log.Fatal(http.ListenAndServe(env.LISTEN_ADDRESS, handler))
}

func configureRoutes(db *sql.DB, router *mux.Router) {
	routes.EnableHealthcheckRoutes(db, router)
	routes.EnableServiceRoutes(db, router)
}

func main() {
	// database connection
	db := config.GetDBConnection()

	// handle routing
	router := mux.NewRouter()
	configureRoutes(db, router)

	// server config
	handler := configureCors(router)
	startServer(handler)

	db.Close()
}
