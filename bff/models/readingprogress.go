package models

type ReadingProgress struct {
	ID          uint   `json:"id"`
	UserId      int    `json:"userId"`
	BookId      int    `json:"bookId"`
	PagesNumber int    `json:"pagesNumber"`
	CreatedAt   string `json:"createdAt"`
}

type CreateReadingProgress struct {
	UserId      int `json:"userId"`
	BookId      int `json:"bookId"`
	PagesNumber int `json:"pagesNumber"`
}
