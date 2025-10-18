package model

type TokenEmail struct {
	Email string `json:"email"`
}

type Token struct {
	User     *User
	IsOnline bool `json:"is_online"`
}
