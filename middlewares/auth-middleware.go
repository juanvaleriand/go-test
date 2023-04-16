package middlewares

import (
	"net/http"
	"os"

	"github.com/juanvaleriand/go-test/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")

		if authToken == "" {
			utils.Respond(w, http.StatusUnauthorized, "Authorization token missing")
			return
		}

		if isValidToken(authToken) {
			next.ServeHTTP(w, r)
			return
		}

		utils.Respond(w, http.StatusUnauthorized, "Invalid authorization token")
	})
}

func isValidToken(token string) bool {
	if token == os.Getenv("TOKEN") {
		return true
	}

	return false
}
