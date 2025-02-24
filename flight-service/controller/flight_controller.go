package controller

import (
	"flight-service/interfaces"
	"flight-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12"
)

type flightController struct {
	flightService interfaces.FlightCRUDService
}

// CreateFlight implements interfaces.FlightCRUDController.
func (f *flightController) CreateFlight(ctx *iris.Context) {
	var flight models.FlightDTO
	if err := ctx.ReadJSON(&flight); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString("Invalid request body")
		return
	}
}

// DeleteFlight implements interfaces.FlightCRUDController.
func (f *flightController) DeleteFlight(ctx *gin.Context) {
	panic("unimplemented")
}

// GetFlight implements interfaces.FlightCRUDController.
func (f *flightController) GetFlight(ctx *gin.Context) {
	panic("unimplemented")
}

// UpdateFlight implements interfaces.FlightCRUDController.
func (f *flightController) UpdateFlight(ctx *gin.Context) {
	panic("unimplemented")
}

func FlightController(flightService interfaces.FlightCRUDService) interfaces.FlightCRUDController {
	return &flightController{
		flightService: flightService,
	}
}
