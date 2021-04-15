package users

import (
	"errors"

	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
)

func PutUser(u models.User, user string) (bool, error) {
	db, _ := db.GetDB()
	t, _, exist := CheckUserExists(u.Username)
	if exist {
		if user != t.Username {
			return false, errors.New("user already exists")
		}
	}
	_, err := db.Exec(`UPDATE users SET username=?, email=?, password=?, name=?, last_name=?, birthday=? WHERE id=?;`, u.Username, u.Email, u.Password, u.Name, u.LastName, u.Birthday, u.Id)
	if err != nil {
		panic(err)
	}
	return true, nil
}
