package env

// import "os"

var (
	DATABASE_URL    = "postgres://tinkerer:password@localhost:4666/tinker?sslmode=disable"
	ALLOWED_ORIGINS = "*"
	LISTEN_ADDRESS  = ":8010"

	ClientId     = "citural-microservice"
	ClientSecret = "3aUXKqK9ac1rdoHZSsZ6XN9ZvefSSOjn"
	Realm        = "citural"
	Hostname     = "http://localhost:8800/"
)

func LoadEnv() {
	// ClientId = os.Getenv("CLIENT_ID")
	// ClientSecret = os.Getenv("CLIENT_SECRET")
	// Realm = os.Getenv("REALM")
	// Hostname = os.Getenv("HOST")
}
