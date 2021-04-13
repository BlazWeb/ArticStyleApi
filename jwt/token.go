package jwt

import (
	"time"

	"github.com/Artic-Dev/ArticStyleApi-GO/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func TokenJWT(t models.User) (string, error) {
	pin := []byte("ArticSolutionsMandaCojooones")

	json := jwt.MapClaims{
		"id":              t.Id,
		"username":        t.Username,
		"password":        t.Password,
		"email":           t.Email,
		"name":            t.Name,
		"lastname":        t.LastName,
		"img":             t.Img,
		"date_registered": t.DateRegistered,
		"rank":            t.Rank,
		"expire":          time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, json)

	tokenStr, err := token.SignedString(pin)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
