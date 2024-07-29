package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/niemet0502/shirabe/bff/models"
	"github.com/niemet0502/shirabe/bff/services"
)

type ShelvesHandler struct {
	svc *services.ShelvesService
}

func NewShelvesHandler(svc *services.ShelvesService) *ShelvesHandler {
	return &ShelvesHandler{svc}
}

func (h *ShelvesHandler) GetShelves(w http.ResponseWriter, r *http.Request) {
	// TODO get userId from token Authorization
	userId := 2

	result, err := h.svc.GetShelves(userId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (h *ShelvesHandler) GetShelf(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	book, err := h.svc.GetShelf(id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(book)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *ShelvesHandler) CreateShelf(w http.ResponseWriter, r *http.Request) {
	// TODO get the userId from the token
	userId := 50

	var createShelf models.CreateShelf

	err := json.NewDecoder(r.Body).Decode(&createShelf)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createShelf.UserId = userId

	result, err := h.svc.CreateShelf(createShelf)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *ShelvesHandler) UpdateShelf(w http.ResponseWriter, r *http.Request) {
	var toUpdate models.Shelf
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	err := json.NewDecoder(r.Body).Decode(&toUpdate)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	toUpdate.ID = id

	result, er := h.svc.UpdateShelf(toUpdate)

	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(result)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *ShelvesHandler) RemoveShelf(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	_, er := h.svc.RemoveShelf(id)

	if er != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
