package services

import (
	"flight-service/interfaces"
	"flight-service/models"

	"go.uber.org/zap"
)

type flightService struct {
	flightRepository interfaces.FlightRepository
	logger           *zap.Logger
}

func flightMapping(flightDto models.FlightDTO) models.Flight {
	return models.Flight{
		FlightNumber:  flightDto.FlightNumber,
		Destination:   flightDto.Destination,
		Airline:       flightDto.Airline,
		Origin:        flightDto.Origin,
		DepartureTime: flightDto.DepartureTime,
		ArrivalTime:   flightDto.ArrivalTime,
	}
}

func flightResponseMapping(flight models.Flight) models.FlightResponse {
	return models.FlightResponse{
		ID:        string(rune(flight.Model.ID)),
		CreatedAt: flight.CreatedAt,
		UpdatedAt: flight.UpdatedAt,
		DelatedAt: &flight.DeletedAt.Time,
		FlightDTO: models.FlightDTO{
			FlightNumber:  flight.FlightNumber,
			Destination:   flight.Destination,
			Airline:       flight.Airline,
			Origin:        flight.Origin,
			DepartureTime: flight.DepartureTime,
			ArrivalTime:   flight.ArrivalTime,
		},
	}
}

// CreateFlight implements interfaces.FlightCRUDService.
func (f *flightService) CreateFlight(flight models.FlightDTO) (models.FlightResponse, error) {
	flightModel := flightMapping(flight)
	createdFlight, err := f.flightRepository.Create(&flightModel)
	if err != nil {
		f.logger.Error("Failed to create flight :: services")
		return models.FlightResponse{}, err
	}
	f.logger.Info("Flight created successfully :: services")
	return flightResponseMapping(*createdFlight), nil
}

// DeleteFlight implements interfaces.FlightCRUDService.
func (f *flightService) DeleteFlight(flightID uint) error {
	err := f.flightRepository.Delete(flightID)
	if err != nil {
		f.logger.Error("Failed to delete flight :: services")
		return err
	}
	f.logger.Info("Flight deleted successfully :: services")
	return nil
}

// GetFlight implements interfaces.FlightCRUDService.
func (f *flightService) GetFlight(flightID uint) (models.FlightResponse, error) {
	flight, err := f.flightRepository.Get(flightID)
	if err != nil {
		f.logger.Error("Failed to get flight :: services")
		return models.FlightResponse{}, err
	}
	f.logger.Info("Flight retrieved successfully :: services")
	return flightResponseMapping(*flight), nil
}

// UpdateFlight implements interfaces.FlightCRUDService.
func (f *flightService) UpdateFlight(flightID uint, flight models.FlightDTO) (models.FlightResponse, error) {
	flightModel := flightMapping(flight)
	updatedFlight, err := f.flightRepository.Update(flightID, &flightModel)
	if err != nil {
		f.logger.Error("Failed to update flight :: services")
		return models.FlightResponse{}, err
	}
	f.logger.Info("Flight updated successfully :: services")
	return flightResponseMapping(*updatedFlight), nil
}

func NewFlightService(repo interfaces.FlightRepository, logger *zap.Logger) interfaces.FlightCRUDService {
	return &flightService{
		flightRepository: repo,
		logger:           logger,
	}
}
