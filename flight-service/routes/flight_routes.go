package routes

import (
	"flight-service/controller"
	"flight-service/repository"
	"flight-service/services"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func FlightRoutes(router *iris.Application, logger *zap.Logger, db *gorm.DB) {
	flight_repo := repository.NewFlightRepository(db, logger)
	flight_service := services.NewFlightService(flight_repo, logger)
	flight_controller := controller.FlightController(flight_service)

	router.Post("/flights", flight_controller.CreateFlight)
	router.Get("/flights/{id:uint}", flight_controller.GetFlight)
	router.Get("/flights/all", flight_controller.GetAllFlights)
	router.Put("/flights/{id:uint}", flight_controller.UpdateFlight)
	router.Delete("/flights/{id:uint}", flight_controller.DeleteFlight)
}
