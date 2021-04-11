package inserts

import (
	"time"

	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/helpers"
)

func InsertRegisterUser(user string, email string, name string, last_name string, password string) (bool, error) {
	register := time.Now()
	db, err := db.GetDB()
	if err != nil {
		return false, err
	}
	// Encriptación de la contraseña
	password, err = helpers.EncryptPassword(password)
	if err != nil {
		return false, err
	}

	_, err = db.Exec("INSERT INTO users (username, email, name, last_name, password, date_registered) VALUES (?, ?, ?, ?, ?, ?)", user, email, name, last_name, password, register)
	return true, err
}
