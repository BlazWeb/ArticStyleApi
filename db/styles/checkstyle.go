package styles

import (
	"database/sql"
	"errors"

	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
)

func CheckStyleID(id int64) (models.Style, error) {
	db, err := db.GetDB()
	var s models.Style
	if err != nil {
		return s, err
	}

	// Consulta a la base de datos si existe un estilo con ese id correspondiente
	row := db.QueryRow("SELECT id, title, author, label, created FROM styles WHERE id = ?", id)
	err = row.Scan(&s.Id, &s.Title, &s.Author, &s.Label, &s.Registrado)
	if err != nil {
		return s, err
	}
	return s, nil
}

func CheckStyleExist(id int64) error {
	db, err := db.GetDB()
	var s models.Style
	if err != nil {
		return err
	}

	// Consulta a la base de datos si existe un estilo con ese id correspondiente
	row := db.QueryRow("SELECT id FROM styles WHERE id = ?", id)
	err = row.Scan(&s.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("No se ha encotrado el estilo")
		}
		return err
	}
	return nil
}

func GetStyleAll(id int64) ([]models.Style, error) {
	db, err := db.GetDB()
	s := []models.Style{}
	if err != nil {
		return s, err
	}

	// Consulta a la base de datos si existe un estilo con ese id correspondiente
	rows, err := db.Query("SELECT id, title, author, label, created FROM styles WHERE author = ?", id)
	if err != nil {
		return s, err
	}
	//err = row.Scan(&s.Id, &s.Title, &s.Author, &s.Label, &s.Registrado)
	for rows.Next() {
		var t models.Style
		err = rows.Scan(&t.Id, &t.Title, &t.Author, &t.Label, &t.Registrado)
		if err != nil {
			return s, err
		}
		s = append(s, t)
	}
	return s, nil
}

func GetStyleSavedAll(id int64) ([]models.StyleSaved, error) {
	db, err := db.GetDB()
	s := []models.StyleSaved{}
	if err != nil {
		return s, err
	}

	// Consulta a la base de datos si existe un estilo con ese id correspondiente
	rows, err := db.Query("SELECT id, style_id, user_id, love FROM styles_save WHERE user_id=?", id)
	if err != nil {
		return s, err
	}
	//err = row.Scan(&s.Id, &s.Title, &s.Author, &s.Label, &s.Registrado)
	for rows.Next() {
		var t models.StyleSaved
		err = rows.Scan(&t.Id, &t.Style, &t.User, &t.Love)
		if err != nil {
			if err == sql.ErrNoRows {
				return s, errors.New("No tiene estilos guardados")
			}
			return s, err
		}
		s = append(s, t)
	}
	return s, nil
}
