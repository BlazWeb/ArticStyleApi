package styles

import (
	"errors"

	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/db/users"
)

func SaveStyleDB(iduser int64, idstyle int64) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	err = users.CheckUserID(iduser)
	if err != nil {
		return errors.New("El usuario no existe")
	}
	err = CheckStyleExist(idstyle)
	if err != nil {
		return errors.New("El estilo no existe")
	}
	_, err = db.Exec("INSERT INTO style_save (style_id, user_id) VALUES (?, ?)", iduser, idstyle)
	if err != nil {
		return err
	}
	return nil
}
