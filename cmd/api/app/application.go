package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pos_system/cmd/api/app/dependence"
)

func New() {
	handlerContainer := dependence.NewWire()
	routes(&handlerContainer)
}

func routes(container *dependence.HandlerContainer) {
	mux := mux.NewRouter()

	mux.HandleFunc("/locations/countries", container.HandlerLocation.GetAllCountries).Methods("GET")
	mux.HandleFunc("/locations/save", container.HandlerLocation.SaveCountries).Methods("POST")
	mux.HandleFunc("/locations/update", container.HandlerLocation.UpdateCountries).Methods("PUT")

	log.Print("Run Server: localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
