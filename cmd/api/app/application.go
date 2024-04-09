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

	mux.HandleFunc("/countries", container.HandlerLocation.GetAllContries).Methods("GET")

	log.Print("Run Server: localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
