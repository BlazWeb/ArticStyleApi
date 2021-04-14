package posts

import (
	"time"

	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
)

func InsertPost(p *models.Post) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	var register = time.Now()

	_, err = db.Exec("INSERT INTO posts (profile, author, content, time) VALUES (?, ?, ?, ?)", p.Profile, p.Author, p.Content, register)
	if err != nil {
		return err
	}
	return nil
}
