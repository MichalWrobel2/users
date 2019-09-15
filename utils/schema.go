package utils

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}
