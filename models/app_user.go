package models

type AppUser struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	Password        string `json:"-"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	Token           string `json:"token"`
	IsBlocked       bool   `json:"is_blocked"`
	IsAdministrator bool   `json:"is_administrator"`
}
