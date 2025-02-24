package interfaces

import "flight-service/models"

type IRepository[T any] interface {
	Create(entity *T) (*T, error)
	Get(id uint) (*T, error)
	GetAll() ([]*T, error)
	Update(id uint, entity *T) (*T, error)
	Delete(id uint) error
}

type FlightRepository interface {
	IRepository[models.Flight]
}
