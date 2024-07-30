package models

type BookShelf struct {
	ID      int `json:"id"`
	ShelfId int `json:"shelfId"`
	BookId  int `json:"bookId"`
}
