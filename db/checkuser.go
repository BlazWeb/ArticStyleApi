package db

import "github.com/Artic-Dev/ArticStyleApi-GO/models"

func CheckUserEmail(email string) (string, error) {
	db, err := GetDB()
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
	db, err := GetDB()
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
