package models

type User struct {
	Id             int64  `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	DateRegistered string `json:"date_registered"`
	Birthday       string `json:"birthday"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Rank           int    `json:"rank"`
	Img            string `json:"img"`
}
