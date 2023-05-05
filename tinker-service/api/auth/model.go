package auth

import (
	"os"

	"github.com/Nerzal/gocloak/v13"
)

var (
	// clientId     = os.Getenv("CLIENT_ID")
	// clientSecret = os.Getenv("CLIENT_SECRET")
	// realm        = os.Getenv("REALM")
	// hostname     = os.Getenv("HOST")
	clientId     string
	clientSecret string
	realm        string
	hostname     string
)
var client gocloak.GoCloak

func InitializeOauthServer() {
	clientId = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	realm = os.Getenv("REALM")
	hostname = os.Getenv("HOST")
	
	client = *gocloak.NewClient(hostname)
}
