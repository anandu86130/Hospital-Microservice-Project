package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email       string
	Role        string
	PayloadHash string
	jwt.StandardClaims
}

// GenerateToken will generate for 1 hour with given data
func GenerateToken(key, email string) (string, error) {
	expTime := time.Now().Add(5 * time.Hour).Unix()

	claims := &Claims{
		Email:       email,
		Role:        "admin",
		PayloadHash: AdminhashPayload(email),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
			Subject:   email,
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Use combination of email and key for signing
	signingKey := []byte(key)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString(signingKey)
	if err != nil {
		log.Printf("unable to generate token for admin %v, err: %v", email, err.Error())
		return "", err
	}
	return signedToken, nil
}

func AdminhashPayload(email string) string {
	h := sha256.New()
	h.Write([]byte(email))
	return hex.EncodeToString(h.Sum(nil))
}

// type Claims struct {
// 	Email string
// 	Role  string
// 	jwt.StandardClaims
// }

// // GenerateToken will generate for 1 hour with given data
// func GenerateToken(key, email string) (string, error) {
// 	expTime := time.Now().Add(1 * time.Hour).Unix()

// 	claims := &Claims{
// 		Email: email,
// 		Role:  "admin",
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expTime,
// 			Subject:   email,
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 	}
// 	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedToken, err := jwtToken.SignedString([]byte(key))
// 	if err != nil {
// 		log.Printf("unable to generate token for admin %v, err: %v", email, err.Error())
// 		return "", err
// 	}
// 	return signedToken, nil
// }
