package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type UserJSON map[string]interface{}

func (u *UserJSON) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &u)
		return nil
	case string:
		json.Unmarshal([]byte(v), &u)
		return nil
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
func (u *UserJSON) Value() (driver.Value, error) {
	return json.Marshal(u)
}

type User struct {
	Id             int64     `json:"id,omitempty"`
	Username       string    `json:"username,omitempty"`
	Name           string    `json:"name,omitempty"`
	LastName       string    `json:"last_name,omitempty"`
	DateRegistered string    `json:"date_registered,omitempty"`
	Birthday       time.Time `json:"birthday,omitempty"`
	Password       string    `json:"password,omitempty"`
	Email          string    `json:"email,omitempty"`
	Rank           int       `json:"rank,omitempty"`
	IP             string    `json:"ip,omitempty"`
	LastIP         string    `json:"last_ip,omitempty"`
	Avatar         string    `json:"avatar,omitempty"`
	Banner         string    `json:"banner,omitempty"`
	// Img struct {
	// 	Avatar string `json:"avatar"`
	// 	Banner string  `json:"banner"`
	// }	`json:"img"`
}

type UserFollower struct {
	Id       int64 `json:"id,omitempty"`
	Author   int   `json:"author,omitempty"`
	Follower int   `json:"Follower,omitempty"`
}

/*RespuestaLogin tiene el token que se devuelve con el login */
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
