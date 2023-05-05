package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/saint-rivers/tinker/env"
)

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

func Protect(next http.Handler) http.Handler {
	// env.LoadEnv()
	client := InitializeOauthServer()

	handler := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) < 1 {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(UnauthorizedError())
			return
		}

		accessToken := strings.Split(authHeader, " ")[1]
		rptResult, err := client.RetrospectToken(r.Context(), accessToken, env.ClientId, env.ClientSecret, env.Realm)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(BadRequestError(err.Error()))
			return
		}
		isTokenValid := *rptResult.Active
		if !isTokenValid {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(UnauthorizedError())
			return
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handler)
}
