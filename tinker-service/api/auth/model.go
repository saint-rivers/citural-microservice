package auth

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/saint-rivers/tinker/env"
)

var client gocloak.GoCloak

func InitializeOauthServer() *gocloak.GoCloak {
	client = *gocloak.NewClient(env.Hostname)
	return &client
}
