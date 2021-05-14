package models

import "github.com/dgrijalva/jwt-go"

//User struct declaration
type User struct {
	Username string `json:"username", db:"username"`
	Password string `json:"password", db:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
