package configs

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	Username string
}
