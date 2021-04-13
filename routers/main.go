package routers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Artic-Dev/ArticStyleApi-GO/db/users"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
	jwt "github.com/dgrijalva/jwt-go"
)

type sendmessage struct {
	Message string `json:"message"`
	Status  bool   `json:"success"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Rest Api ArticStylses. Att: ArticDev")
}

var User string
var IDUser string

func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("ArticSolutionsMandaCojooones")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, _, status := users.CheckUserExists(claims.User)
		if status {
			User = claims.User
			IDUser = claims.ID
		}
		return claims, status, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}
	return claims, false, string(""), err
}
