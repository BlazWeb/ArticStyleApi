package users

import (
	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
)

func CheckUserEmail(email string) (string, error) {
	db, err := db.GetDB()
	var u models.User
	if err != nil {
		return "0", err
	}

	row := db.QueryRow("SELECT email FROM users WHERE email = ?", email)
	err = row.Scan(&u.Email)
	if err != nil {
		return "1", nil
	}
	return "0", nil
}

func CheckUserName(user string) (string, error) {
	db, err := db.GetDB()
	var u models.User
	if err != nil {
		return "", err
	}

	row := db.QueryRow("SELECT username FROM users WHERE username = ?", user)
	err = row.Scan(&u.Username)
	if err != nil {
		return "valid", nil
	}
	return "no_valid", nil
}

func CheckUserID(id int64) (models.User, error) {
	db, err := db.GetDB()
	var u models.User
	if err != nil {
		return u, err
	}
	row := db.QueryRow("SELECT id, username, password, email, name, last_name, rank, img, birthday, date_registered  FROM users WHERE id = ?", id)
	err = row.Scan(&u.Id, &u.Username, &u.Password, &u.Email, &u.Name, &u.LastName, &u.Rank, &u.Img, &u.Birthday, &u.DateRegistered)
	if err != nil {
		return u, err
	}
	return u, nil
}
