package middlewares

import "github.com/golang-jwt/jwt/v5"

type ClaimsRefreshToken struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type ClaimsAccessToken struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
