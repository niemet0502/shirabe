package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/niemet0502/shirabe/bff/models"
	"github.com/niemet0502/shirabe/bff/services"
)

type ReadingProgressHandler struct {
	svc *services.ReadingProgressService
}

func NewReadingProgressHandler(svc *services.ReadingProgressService) *ReadingProgressHandler {
	return &ReadingProgressHandler{svc}
}

func (h *ReadingProgressHandler) Create(w http.ResponseWriter, r *http.Request) {
	var createProgress models.CreateReadingProgress

	err := json.NewDecoder(r.Body).Decode(&createProgress)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := h.svc.Create(createProgress)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *ReadingProgressHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["bookId"])

	progress, err := h.svc.GetByBookId(id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(progress)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
