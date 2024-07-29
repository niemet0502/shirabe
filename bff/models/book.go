package models

type Book struct {
	ID              uint   `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Genre           string `json:"genre"`
	Status          int    `json:"status"`
	TotalPages      int    `json:"totalPages"`
	ReadingProgress int    `json:"readingProgress"`
	UserId          int    `json:"userId"`
	Cover           string `json:"cover"`
	Description     string `json:"description"`
	PublicatedAt    string `json:"publicatedAt"`
	StartReadingAt  string `json:"startReadingAt"`
	FinishReadingAt string `json:"finishReadingAt"`
}

type CreateBook struct {
	Title        string `json:"title"`
	Author       string `json:"author"`
	Genre        string `json:"genre"`
	TotalPages   int    `json:"totalPages"`
	UserId       int    `json:"userId"`
	Cover        string `json:"cover"`
	Description  string `json:"description"`
	PublicatedAt string `json:"publicatedAt"`
}
