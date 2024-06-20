package model

import "github.com/dgrijalva/jwt-go"

type (
	JWTClaim struct {
		jwt.StandardClaims
		Username string `json:"username"`
		Roles    string `json:"role"`
	}
)
