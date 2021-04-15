package users

import (
	"time"

	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/helpers"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
)

func InsertRegisterUser(u models.User) (bool, error) {
	register := time.Now()
	db, err := db.GetDB()
	if err != nil {
		return false, err
	}
	// Add Gravatar
	avatar := helpers.RegisterAvatar(u.Email)

	// Encriptación de la contraseña
	password, err := helpers.EncryptPassword(u.Password)
	if err != nil {
		return false, err
	}

	_, err = db.Exec("INSERT INTO users (username, email, name, last_name, password, date_registered, avatar, birthday, ip, last_ip) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", u.Username, u.Email, u.Name, u.LastName, password, register, avatar, u.Birthday, u.IP, u.IP)
	return true, err
}
