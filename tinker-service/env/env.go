package env

// import "os"

var (
	DATABASE_URL    = "postgres://tinkerer:password@localhost:4666/tinker?sslmode=disable"
	ALLOWED_ORIGINS = "http://localhost:4200"
	LISTEN_ADDRESS  = ":8010"

	ClientId     = "citural-microservice"
	ClientSecret = "wV66R0Q2Y1Q1SERS5bty4Igq41x6HGYw"
	Realm        = "citural"
	Hostname     = "http://localhost:8800/"
)

func LoadEnv() {
	// ClientId = os.Getenv("CLIENT_ID")
	// ClientSecret = os.Getenv("CLIENT_SECRET")
	// Realm = os.Getenv("REALM")
	// Hostname = os.Getenv("HOST")
}
