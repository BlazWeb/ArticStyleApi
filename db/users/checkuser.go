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
	s := `SELECT * FROM users WHERE id=?;`
	row := db.QueryRow(s, id)
	err = row.Scan(&u.Id, &u.Email, &u.Username, &u.Password, &u.Name, &u.LastName, &u.DateRegistered, &u.Birthday, &u.Rank, &u.IP, &u.LastIP, &u.Avatar, &u.Banner)
	//err = row.Scan(&u.Id, &u.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("No user found")
		}
		return u, err
	}
	return u, nil
}

func CheckUserFollowers(id int64) ([]models.UserFollower, error) {
	db, err := db.GetDB()
	u := []models.UserFollower{}
	if err != nil {
		return u, err
	}
	rows, err := db.Query("SELECT id, author_id, follower_id FROM users_followers WHERE author_id = ?", id)
	if err != nil {
		return u, err
	}
	for rows.Next() {
		var t models.UserFollower
		err = rows.Scan(&t.Id, &t.Author, &t.Follower)
		if err != nil {
			return u, err
		}
		u = append(u, t)
	}
	return u, err
}

func CheckUserExists(user string) (models.User, error, bool) {
	db, err := db.GetDB()
	var u models.User
	if err != nil {
		return u, err, false
	}
	row := db.QueryRow("SELECT * FROM users WHERE username = ?", user)
	err = row.Scan(&u.Id, &u.Email, &u.Username, &u.Password, &u.Name, &u.LastName, &u.DateRegistered, &u.Birthday, &u.Rank, &u.IP, &u.LastIP, &u.Avatar, &u.Banner)
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
