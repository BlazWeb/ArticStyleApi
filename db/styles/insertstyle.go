package styles

import (
	"time"

	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/helpers"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
)

func InsertRegisterStyle(author int64, title string) (models.Style, error) {
	db, err := db.GetDB()
	var t models.Style
	if err != nil {
		return t, err
	}
	sl, err := db.Exec("INSERT INTO styles (author, title, created) VALUES (?, ?, ?)", author, title, time.Now())
	if err != nil {
		return t, err
	}
	sol, err := sl.LastInsertId()
	//row, err := db.Query("SELECT MAX(id) FROM styles WHERE author = ?", author)
	//err = row.Scan(&t.Id)
	if err != nil {
		return t, err
	}
	t.Id = sol
	res, err := helpers.EncryptHash(string(t.Id))
	_, err = db.Exec("UPDATE styles SET label = ? WHERE id = ?", res, t.Id)
	if err != nil {
		return t, err
	}
	row := db.QueryRow("SELECT label FROM styles WHERE id = ?", t.Id)
	err = row.Scan(&t.Label)
	if err != nil {
		return t, err
	}
	return t, nil
}
