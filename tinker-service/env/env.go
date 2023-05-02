package env

var (
	// DATABASE_URL    = "host=ep-tiny-star-976595.ap-southeast-1.aws.neon.tech user=eam.dayan password=gQY3MG9rkOsL dbname=tinker sslmode=disable"
	// DATABASE_URL    = "host=postgres://localhost:4666 user=tinkerer password=password dbname=tinker sslmode=disable"
	DATABASE_URL    = "postgres://tinkerer:password@localhost:4666/tinker?sslmode=disable"
	ALLOWED_ORIGINS = "*"
	LISTEN_ADDRESS  = ":8010"
)
