package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/niemet0502/shirabe/bff/services"
)

type BookShelfHandler struct {
	svc *services.BookShelfService
}

func NewBookShelfHandler(svc *services.BookShelfService) *BookShelfHandler {
	return &BookShelfHandler{svc}
}

func (h *BookShelfHandler) GetBooksByShelf(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["shelfId"])

	res, err := h.svc.GetBooksByShelf(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, _ := json.Marshal(res)

	w.WriteHeader(http.StatusAccepted)
	w.Write(json)
}
func (h *BookShelfHandler) AddBookToShelf(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	shelfId, _ := strconv.Atoi(vars["shelfId"])
	bookId, _ := strconv.Atoi(vars["bookId"])

	res, err := h.svc.AddBookToShelf(shelfId, bookId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, _ := json.Marshal(res)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(json)
}
func (h *BookShelfHandler) RemoveBookFromShelf(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	shelfId, _ := strconv.Atoi(vars["shelfId"])
	bookId, _ := strconv.Atoi(vars["bookId"])

	_, err := h.svc.RemoveBookFromShel(shelfId, bookId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// func (h *BookShelfHandler) (w http.WriteResponse, r *http.Request){}
