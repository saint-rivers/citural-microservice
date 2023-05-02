package database

import (
	"database/sql"

	request "github.com/saint-rivers/tinker/common"
)

func InsertDatabaseRequest(db *sql.DB, r *request.DatabaseServiceRequest) *sql.Row {
	err := db.QueryRow(
		`INSERT INTO app_service (container_name, database_vendor, default_database, db_username, db_password)
		VALUES ($1, $2, $3, $4, $5);
		`,
		r.ContainerName, r.DatabaseVendor, r.DefaultDatabase, r.UserName, r.Password,
	)
	return err
}
