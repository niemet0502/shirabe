package models

type Book struct {
	ID              int
	Title           string
	Author          string
	Genre           string
	Status          int
	TotalPages      int
	ReadingProgress int
	UserId          int
}

type CreateBook struct {
	Title      string
	Author     string
	Genre      string
	TotalPages int
	UserId     int
}
