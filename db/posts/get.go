package posts

import (
	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
)

func GetPostsProfile(id int64) ([]models.Post, error) {
	db, err := db.GetDB()
	posts := []models.Post{}
	if err != nil {
		return posts, err
	}

	// Consulta a la base de datos si existe un estilo con ese id correspondiente
	rows, err := db.Query("SELECT id, author, profile, content, likes FROM posts WHERE profile = ?", id)
	if err != nil {
		return posts, err
	}
	//err = row.Scan(&s.Id, &s.Title, &s.Author, &s.Label, &s.Registrado)
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.Id, &post.Author, &post.Profile, &post.Content, &post.Likes)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
