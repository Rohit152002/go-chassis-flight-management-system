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
}
