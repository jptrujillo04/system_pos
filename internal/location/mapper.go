package location

import (
	"encoding/json"
	"errors"
	"net/http"
)

func MapRequestToCountryRequest(r *http.Request, output interface{}) error {
	err := json.NewDecoder(r.Body).Decode(output)
	if err != nil {
		return errors.New("Error parsing request country body: " + err.Error())
	}
	return nil
}

func MapCountryRequestToModelCountry(request CountryRequest) Country {
	return Country{
		Name:   request.Name,
		Status: request.Status,
	}
}
