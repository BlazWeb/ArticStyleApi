package models

type User struct {
	Id             int64  `json:"id,omitempty"`
	Username       string `json:"username,omitempty"`
	Name           string `json:"name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	DateRegistered string `json:"date_registered,omitempty"`
	Birthday       string `json:"birthday,omitempty"`
	Password       string `json:"password,omitempty"`
	Email          string `json:"email,omitempty"`
	Rank           int    `json:"rank,omitempty"`
	Img            string `json:"img,omitempty"`
}

/*RespuestaLogin tiene el token que se devuelve con el login */
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
