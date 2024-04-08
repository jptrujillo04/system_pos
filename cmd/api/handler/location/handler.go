package location

import (
	"net/http"
	"pos_system/internal/location"
)

type Handler struct {
	UseCaseLocation location.UseCase
}

func NewHandler(useCaseLocation location.UseCase) Handler {
	return Handler{useCaseLocation}
}

func (h Handler) GetAllContries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message: Holamundo"))
}
