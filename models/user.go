package models

import "time"

type User struct {
	Id             int64     `json:"id"`
	Username       string    `json:"username"`
	Name           string    `json:"name"`
	LastName       string    `json:"last_name"`
	DateRegistered time.Time `json:"date_registered"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	Rank           int       `json:"rank"`
	Img            string    `json:"img"`
}
