package styles

import (
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
