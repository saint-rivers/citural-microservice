package auth

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/saint-rivers/tinker/env"
)

var Client gocloak.GoCloak

func InitializeOauthServer() {
	Client = *gocloak.NewClient(env.Hostname)
}
