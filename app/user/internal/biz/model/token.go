package model

type TokenVerityCodeAccount struct {
	Account string `json:"account"`
}

type Token struct {
	User     *User
	IsOnline bool `json:"is_online"`
}
