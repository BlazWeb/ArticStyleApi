package users

import (
	"github.com/Artic-Dev/ArticStyleApi-GO/db"
)

func DelUser(id int64) error {
	err := CheckUser(id)
	if err != nil {
		return err
	}
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	return nil
}
