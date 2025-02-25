package interfaces

import (
	"github.com/kataras/iris/v12"
)

type FlightCRUDController interface {
	CreateFlight(ctx iris.Context)
	GetFlight(ctx iris.Context)
	UpdateFlight(ctx iris.Context)
	DeleteFlight(ctx iris.Context)
	GetAllFlights(ctx iris.Context)
}
