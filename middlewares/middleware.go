package middlewares

import (
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ihksanghazi/api-online-course/databases"
	"github.com/ihksanghazi/api-online-course/models"
)

type contextKey string

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

		if claims, ok := token.Claims.(*ClaimsToken); ok && token.Valid {
			// Jika token valid, lanjutkan ke handler berikutnya
			ctx := context.WithValue(r.Context(), contextKey("id"), claims.Id)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// Jika token tidak valid, berikan respon error
		http.Error(w, err.Error(), http.StatusUnauthorized)
	})
}

func OnlyTeacherAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// mangambil value context id yang di kirim ke middleware sebelumnya
		id := r.Context().Value(contextKey("id")).(string)

		var user models.User
		// cek database dengan id
		if err := databases.DB.First(&user, "id = ?", id).Error; err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// jika role user bukan admin dan teacher kembalikan error
		if user.Role != "admin" && user.Role != "teacher" {
			http.Error(w, "Not Teacher", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func OnlyAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// mangambil value context id yang di kirim ke middleware sebelumnya
		id := r.Context().Value(contextKey("id")).(string)

		var user models.User
		// cek database dengan id
		if err := databases.DB.First(&user, "id = ?", id).Error; err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// jika role user bukan admin kembalikan error
		if user.Role != "admin" {
			http.Error(w, "Not Admin", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
