package auth1

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("mylonelywolf")

type JwtClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken
func GenerateToken(username string, email string) (tokenString string, err error) {
	//Expiration time
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &JwtClaim{
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(secretKey)
	return
}

// Validate token
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return
	}
	//Seperate component of token
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't be parse claims")
		return
	}
	//Time experied of Token
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("Token expired")
		return
	}
	return

}
