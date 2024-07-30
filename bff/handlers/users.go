package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/niemet0502/shirabe/bff/models"
	"github.com/niemet0502/shirabe/bff/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// writeErrorResponse is a helper function to send an error response
func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}

type UserHandler struct {
	svc *services.UserService
}

func NewUserHandler(svc *services.UserService) *UserHandler {
	return &UserHandler{svc}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUser models.CreateUser

	err := json.NewDecoder(r.Body).Decode(&createUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	result, er := h.svc.CreateUser(createUser)

	if er != nil {

		errStatus, _ := status.FromError(er)

		if errStatus.Code() == codes.AlreadyExists {
			writeErrorResponse(w, http.StatusBadRequest, "Email already in use")
			return
		}

		writeErrorResponse(w, http.StatusBadRequest, "An error has occured")
		return
	}

	res, _ := json.Marshal(result)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
