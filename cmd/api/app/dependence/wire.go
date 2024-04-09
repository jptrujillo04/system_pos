package dependence

import (
	"log"
	handlerLocation "pos_system/cmd/api/handler/location"
	"pos_system/internal/location"
)

type HandlerContainer struct {
	HandlerLocation handlerLocation.Handler
}

func NewWire() HandlerContainer {
	dep, err := NewDependencies()
	if err != nil {
		log.Panicf(err.Error())
		return HandlerContainer{}
	}
	useCaseLocation := location.NewUseCaseLocation(dep.LocationRepository)
	return HandlerContainer{
		HandlerLocation: handlerLocation.NewHandler(useCaseLocation),
	}
}
