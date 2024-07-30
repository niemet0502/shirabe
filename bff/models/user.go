package models

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type CreateUser struct {
	Email    string
	Password string
}
