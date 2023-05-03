package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/Nerzal/gocloak"
)

var (
	clientId     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	realm        = os.Getenv("REALM")
	hostname     = os.Getenv("HOST")
)
var client gocloak.GoCloak

func InitializeOauthServer() {
	client = gocloak.NewClient(hostname)
}

func Protect(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) < 1 {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(UnauthorizedError())
			return
		}

		accessToken := strings.Split(authHeader, " ")[1]
		rptResult, err := client.RetrospectToken(accessToken, clientId, clientSecret, realm)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(BadRequestError(err.Error()))
			return
		}
		isTokenValid := rptResult.Active
		if !isTokenValid {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(UnauthorizedError())
			return
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handler)
}
