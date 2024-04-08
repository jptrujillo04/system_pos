package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pos_system/cmd/api/app/dependence"
	"pos_system/internal/config"
	"pos_system/internal/database"
)

func New() {
	mapRoutes()
}

func mapRoutes() {
	dbConfig := config.ReadDBConfig()

	db, err := database.ConnectDB(dbConfig)
	if err != nil {
		log.Println("Error conectando a la base de datos:", err)
	}

	defer db.Close()
	handlerContainer := dependence.NewWire(db)
	routes(&handlerContainer)
}

func routes(container *dependence.HandlerContainer) {
	mux := mux.NewRouter()

	mux.HandleFunc("/countries", container.HandlerLocation.GetAllContries).Methods("GET")

	log.Print("Run Server: localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
