package middlewares

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// mengambil header
		accessToken := r.Header.Get("Access-Token")
		if accessToken == "" {
			http.Error(w, "No Access Token", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(accessToken, &ClaimsToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_JWT_KEY")), nil
		})

		if _, ok := token.Claims.(*ClaimsToken); ok && token.Valid {
			// Jika token valid, lanjutkan ke handler berikutnya
			next.ServeHTTP(w, r)
			return
		}

		// Jika token tidak valid, berikan respon error
		http.Error(w, err.Error(), http.StatusUnauthorized)
	})
}
