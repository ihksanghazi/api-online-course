package middlewares

import "github.com/golang-jwt/jwt/v5"

type ClaimsToken struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}
