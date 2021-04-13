package models

type Style struct {
	Id         int64  `json:"id"`
	Author     int64  `json:"author"`
	Title      string `json:"title,omitempty"`
	Likes      int    `json:"likes,omitempty"`
	Downloads  int    `josn:"download, omitenmpty"`
	Registrado string `json:"registrado"`
	Label      string `json:"label"`
}
