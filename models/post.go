package models

type Post struct {
	Id      int64  `json:"id"`
	Profile int64  `json:"profile,omitempty"`
	Author  int64  `json:"author,omitempty"`
	Content string `json:"content,omitempty"`
	Likes   int    `json:"likes,omitempty"`
}
