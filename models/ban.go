package models

type Ban struct {
	Id       int    `json:"id,omitempty"`
	User     int64  `json:"user,omitempty"`
	Ip       string `json:"ip,omitempty"`
	Reason   string `json:"reason,omitempty"`
	Date     string `json:"date,omitempty"`
	Duration string `json:"duration,omitempty"`
}
