package models

type Shelf struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserId int    `json:"userId"`
	Books  []Book `json:"books"`
}

type CreateShelf struct {
	Name   string `json:"name"`
	UserId int    `json:"userId"`
}
