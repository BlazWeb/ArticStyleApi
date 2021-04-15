package users

import (
	"database/sql"
	"time"

	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
)

func CheckBanIP(ip string) (models.Ban, bool, error) {
	var b models.Ban
	db, err := db.GetDB()
	if err != nil {
		return b, false, err
	}
	time := time.Now()
	row := db.QueryRow(`SELECT * FROM users_bans WHERE ip=? AND duration > ?;`, ip, time)
	err = row.Scan(&b.Id, &b.User, &b.Ip, &b.Reason, &b.Date, &b.Duration)
	if err != nil {
		if err == sql.ErrNoRows {
			return b, false, nil
		}
		return b, false, err
	}
	return b, true, nil
}

// Misma función pero chequeando mediante el usuario
func CheckBanUser(id int64) (models.Ban, bool, error) {
	var b models.Ban
	db, err := db.GetDB()
	if err != nil {
		return b, false, err
	}
	time := time.Now()
	row := db.QueryRow(`SELECT * FROM users_bans WHERE user_id=? AND duration > ?;`, id, time)
	err = row.Scan(&b.Id, &b.User, &b.Ip, &b.Reason, &b.Date, &b.Duration)
	if err != nil {
		if err == sql.ErrNoRows {
			return b, false, nil
		}
		return b, false, err
	}
	return b, true, nil
}

func BanUser(b models.Ban, ip string) (bool, error) {
	db, err := db.GetDB()
	if err != nil {
		return false, err
	}
	date := time.Now()
	_, err = db.Exec(`INSERT INTO users_bans (user_id, ip, reason, time, duration) VALUES (?, ?, ?, ?, ?)`, b.User, ip, b.Reason, date, b.Duration)
	if err != nil {
		return false, err
	}
	return true, nil
}
