package model

import "github.com/dgrijalva/jwt-go"

type AdminClaims struct {
	UserID   uint
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
