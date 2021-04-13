package users

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
	"golang.org/x/crypto/bcrypt"
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
		if err == sql.ErrNoRows {
			return "1", nil
		}
		log.Println(err)
		return "0", err
	}
	return "0", nil
}

func CheckUserNameEspecial(user string) (string, error) {
	db, err := db.GetDB()
	var u models.User
	if err != nil {
		return "", err
	}

	row := db.QueryRow("SELECT username FROM users WHERE username = ?", user)
	err = row.Scan(&u.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "valid", nil
		}
		return "no_valid", err
	}
	return "no_valid", nil
}

func CheckUserLogin(user string, password string) (models.User, error) {
	u, err, _ := CheckUserExists(user)
	if err != nil {
		return u, err
	}
	passwordByte := []byte(password)
	passwordDB := []byte(u.Password)
	err = bcrypt.CompareHashAndPassword(passwordDB, passwordByte)
	if err != nil {
		return u, err
	}
	return u, nil
}

func CheckUserIDEspecial(id int64) (models.User, error) {
	db, err := db.GetDB()
	var u models.User
	if err != nil {
		return u, err
	}
	row := db.QueryRow("SELECT id, username, password, email, name, last_name, img, birthday, date_registered, rank FROM users WHERE id = ?", id)
	err = row.Scan(&u.Id, &u.Username, &u.Password, &u.Email, &u.Name, &u.LastName, &u.Img, &u.Birthday, &u.DateRegistered, &u.Rank)
	//err = row.Scan(&u.Id, &u.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("No user found")
		}
		return u, err
	}
	return u, nil
}

func CheckUserExists(user string) (models.User, error, bool) {
	db, err := db.GetDB()
	var u models.User
	if err != nil {
		return u, err, false
	}
	row := db.QueryRow("SELECT id, username, password, email, name, last_name, img, birthday, date_registered, rank FROM users WHERE username = ?", user)
	err = row.Scan(&u.Id, &u.Username, &u.Password, &u.Email, &u.Name, &u.LastName, &u.Img, &u.Birthday, &u.DateRegistered, &u.Rank)
	if err != nil {
		return u, err, false
	}
	return u, nil, true
}

func CheckUserID(id int64) error {
	db, err := db.GetDB()
	var u models.User
	if err != nil {
		return err
	}
	row := db.QueryRow("SELECT id FROM users WHERE id = ?", id)
	err = row.Scan(&u.Id)
	if err != nil {
		return err
	}
	return nil
}
