package controller

import (
	"flight-service/interfaces"
	"flight-service/models"
	"strconv"

	"github.com/kataras/iris/v12"
)

type flightController struct {
	flightService interfaces.FlightCRUDService
}

// CreateFlight implements interfaces.FlightCRUDController.

// CreateFlightController godoc
// @Summary Create a new flight
// @Description Create a new flight with the input payload
// @Tags flights
// @Accept json
// @Produce json
// @Param user body models.FlightDTO true "Flights"
// @Success 201 {object} models.FlightResponse
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /flights [post]
func (f *flightController) CreateFlight(ctx iris.Context) {
	var flight models.FlightDTO
	if err := ctx.ReadJSON(&flight); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	// Assuming flightService.CreateFlight returns the created flight and an error
	createdFlight, err := f.flightService.CreateFlight(flight)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	ctx.JSON(iris.Map{
		"status": "success",
		"data":   createdFlight,
	})
}

// DeleteFlight implements interfaces.FlightCRUDController.

// DeleteFlightController godoc
// @Summary Delete a flight by ID
// @Description Delete a flight by ID
// @Tags flights
// @Accept json
// @Produce json
// @Param id path int true "Flight ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /flights/{id} [delete]
func (f *flightController) DeleteFlight(ctx iris.Context) {
	id, err := strconv.ParseUint(ctx.Params().Get("id"), 10, 64)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	err = f.flightService.DeleteFlight(uint(id))
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}
	ctx.JSON(iris.Map{
		"status": "success",
		"data":   "Flight deleted successfully",
	})
}

// GetFlight implements interfaces.FlightCRUDController.

// GetFlightController godoc
// @Summary Get a flight by ID
// @Description Get a flight by ID
// @Tags flights
// @Accept json
// @Produce json
// @Param id path int true "Flight ID"
// @Success 200 {object} models.FlightResponse
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /flights/{id} [get]
func (f *flightController) GetFlight(ctx iris.Context) {
	id, err := strconv.ParseUint(ctx.Params().Get("id"), 10, 64)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	flight, err := f.flightService.GetFlight(uint(id))
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}
	ctx.JSON(iris.Map{
		"status": "success",
		"data":   flight,
	})
}

// UpdateFlight implements interfaces.FlightCRUDController.

// UpdateFlightController godoc
// @Summary Update a flight by ID
// @Description Update a flight by ID
// @Tags flights
// @Accept json
// @Produce json
// @Param id path int true "Flight ID"
// @Param user body models.FlightDTO true "Flights"
// @Success 200 {object} models.FlightResponse
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /flights/{id} [put]
func (f *flightController) UpdateFlight(ctx iris.Context) {
	var flight models.FlightDTO
	id, err := strconv.ParseUint(ctx.Params().Get("id"), 10, 64)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	if err := ctx.ReadJSON(&flight); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	updatedFlight, err := f.flightService.UpdateFlight(uint(id), flight)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}
	ctx.JSON(iris.Map{
		"status": "success",
		"data":   updatedFlight,
	})
}

// GetAllFlights implements interfaces.FlightCRUDController.

// GetAllFlightsController godoc
// @Summary Get all flights
// @Description Get all flights
// @Tags flights
// @Accept json
// @Produce json
// @Success 200 {object} []models.FlightResponse
// @Router /flights/all [get]
func (f *flightController) GetAllFlights(ctx iris.Context) {
	var flights []models.FlightResponse
	flights, err := f.flightService.GetAllFlight()
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}
	ctx.JSON(iris.Map{
		"status": "success",
		"data":   flights,
	})

}

func FlightController(flightService interfaces.FlightCRUDService) interfaces.FlightCRUDController {
	return &flightController{
		flightService: flightService,
	}
}
