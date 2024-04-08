package dependence

import (
	"database/sql"
	"log"
	handlerLocation "pos_system/cmd/api/handler/location"
	"pos_system/internal/location"
)

type HandlerContainer struct {
	HandlerLocation handlerLocation.Handler
}

func NewWire(db *sql.DB) HandlerContainer {
	dep, err := NewDependencies(db)
	if err != nil {
		log.Panicf(err.Error())
		return HandlerContainer{}
	}
	useCaseLocation := location.NewUseCaseLocation(dep.LocationRepository)
	return HandlerContainer{
		HandlerLocation: handlerLocation.NewHandler(useCaseLocation),
	}
}
