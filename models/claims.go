package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

/*Claim es la estructura usada para procesar el JWT*/
type Claim struct {
	User string `json:"username"`
	ID   string `json:"id"`
	jwt.StandardClaims
}
