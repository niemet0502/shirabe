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
	userId := 2

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
	userId := 2

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

	var createBook models.CreateBook

	err := json.NewDecoder(r.Body).Decode(&createBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createBook.UserId = userId

	result, err := h.svc.CreateBook(createBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
