package location

import (
	"encoding/json"
	"net/http"
	"pos_system/internal/location"
)

type Handler struct {
	UseCaseLocation location.UseCase
}

func NewHandler(useCaseLocation location.UseCase) Handler {
	return Handler{useCaseLocation}
}

func (h Handler) GetAllCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := h.UseCaseLocation.GetAllCountries(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(countries)
}

func (h Handler) SaveCountries(w http.ResponseWriter, r *http.Request) {
	var countryReq location.CountryRequest
	err := location.MapRequestToCountryRequest(r, &countryReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.UseCaseLocation.SaveCountry(r.Context(), countryReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
