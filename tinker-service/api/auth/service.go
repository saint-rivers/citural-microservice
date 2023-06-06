package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/pkg/errors"
	"github.com/saint-rivers/tinker/env"
)

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

// func ExtractTokenFromHeader(r *http.Request) (string, error) {
// 	authHeader := r.Header.Get("Authorization")
// 	parsedHeader := strings.Split(authHeader, " ")

// 	if parsedHeader[0] == "Bearer" {
// 		return parsedHeader[1]

// 		rptResult, err := Client.RetrospectToken(
// 			r.Context(), tokenString, env.ClientId, env.ClientSecret, env.Realm,
// 		)
// 		if err != nil {
// 			return "", errors.New("unable to retrospect token")
// 		}

// 		isTokenValid := *rptResult.Active
// 		if !isTokenValid {
// 			return "", errors.New("invalid token")
// 		}

// 		return tokenString
// 	}
// }

func extractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	parsedHeader := strings.Split(authHeader, " ")

	if parsedHeader[0] == "Bearer" {
		return parsedHeader[1], nil
	}
	return "", errors.New("unable to parse authorization header")
}

func ExtractKeycloakClaims(r *http.Request) (*gocloak.UserInfo, error) {
	token, _ := extractTokenFromHeader(r)

	return Client.GetUserInfo(r.Context(), token, env.Realm)

	//  Client.RetrospectToken(
	// 	r.Context(), token, env.ClientId, env.ClientSecret, env.Realm,
	// )
}

func Protect(next http.Handler) http.Handler {
	// env.LoadEnv()
	client := Client

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
