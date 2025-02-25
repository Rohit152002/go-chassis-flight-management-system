package interfaces

import "flight-service/models"

type FlightCRUDService interface {
	CreateFlight(flight models.FlightDTO) (models.FlightResponse, error)
	GetFlight(flightID uint) (models.FlightResponse, error)
	UpdateFlight(flightID uint, flight models.FlightDTO) (models.FlightResponse, error)
	DeleteFlight(flightID uint) error
	GetAllFlight() ([]models.FlightResponse, error)
}
