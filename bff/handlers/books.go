package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/niemet0502/shirabe/bff/models"
	"github.com/niemet0502/shirabe/bff/services"

	"github.com/gorilla/mux"
)

type BookHandler struct {
	svc *services.BookService
}

func NewBookHandler(svc *services.BookService) *BookHandler {
	return &BookHandler{svc: svc}
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	book, err := h.svc.GetBook(id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(book)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	// TODO get the userId from the token
	userId := 50

	result, err := h.svc.GetBooks(userId)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *BookHandler) SearchBooks(w http.ResponseWriter, r *http.Request) {
	// TODO get the userId from the token
	userId := 50

	status, _ := strconv.Atoi(r.URL.Query().Get("status"))
	search := r.URL.Query().Get("search")

	result, err := h.svc.BooksSearch(userId, status, search)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	// TODO get the userId from the token
	userId := 50

	err := r.ParseMultipartForm(10 << 20) // Parse form
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	totalPages, _ := strconv.Atoi(r.FormValue("totalPages"))

	createBook := models.CreateBook{
		Title:       r.FormValue("title"),
		Author:      r.FormValue("author"),
		Genre:       r.FormValue("genre"),
		TotalPages:  totalPages,
		UserId:      userId,
		Description: r.FormValue("description"),
	}

	file, handler, err := r.FormFile("cover")
	if err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	createBook.UserId = userId

	result, err := h.svc.CreateBook(createBook, file, handler.Filename)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var toUpdate models.Book
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	err := json.NewDecoder(r.Body).Decode(&toUpdate)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	toUpdate.ID = uint(id)

	result, er := h.svc.UpdateBook(toUpdate)

	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(result)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
