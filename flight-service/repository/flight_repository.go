package repository

import (
	interfaces "flight-service/Interfaces"
	"flight-service/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type flightRepository struct {
	db     *gorm.DB
	Logger *zap.Logger
}

// Create implements interfaces.FlightRepository.
func (f *flightRepository) Create(flight *models.Flight) (*models.Flight, error) {
	if err := f.db.Create(flight).Error; err != nil {
		f.Logger.Info("Failed to create  :: repositories")
		return nil, err
	}
	f.Logger.Info("Flight created successfully :: repositories")
	return flight, nil
}

// Delete implements interfaces.FlightRepository.
func (f *flightRepository) Delete(id uint) error {
	if err := f.db.Delete(&models.Flight{}, id).Error; err != nil {
		f.Logger.Error("Failed to delete :: repositories")
		return err
	}
	f.Logger.Info("Flight deleted successfully :: repositories")
	return nil
}

// Get implements interfaces.FlightRepository.
func (f *flightRepository) Get(id uint) (*models.Flight, error) {
	var flight models.Flight
	if err := f.db.First(&flight, id).Error; err != nil {
		f.Logger.Error("Failed to get :: repositories")
		return nil, err
	}
	f.Logger.Info("Flight retrieved successfully :: repositories")
	return &flight, nil
}

// GetAll implements interfaces.FlightRepository.
func (f *flightRepository) GetAll() ([]*models.Flight, error) {
	var flights []*models.Flight
	if err := f.db.Find(&flights).Error; err != nil {
		f.Logger.Error("Failed to get all :: repositories")
		return nil, err
	}
	f.Logger.Info("All flights retrieved successfully :: repositories")
	return flights, nil
}

// Update implements interfaces.FlightRepository.
func (f *flightRepository) Update(id uint, entity *models.Flight) (*models.Flight, error) {
	if err := f.db.Model(&models.Flight{}).Where("id = ?", id).Updates(entity).Error; err != nil {
		f.Logger.Error("Failed to update :: repositories")
		return nil, err
	}
	f.Logger.Info("Flight updated successfully :: repositories")
	return entity, nil
}

func NewFlightRepository(db *gorm.DB, logger *zap.Logger) interfaces.FlightRepository {
	return &flightRepository{db: db, Logger: logger}
}
