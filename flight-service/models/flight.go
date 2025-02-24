package models

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	FlightNumber  string `json:"flight_number" binding:"required"`
	Airline       string `json:"airline" binding:"required"`
	Origin        string `json:"origin" binding:"required"`
	Destination   string `json:"destination" binding:"required"`
	DepartureTime string `json:"departure_time" binding:"required"`
	ArrivalTime   string `json:"arrival_time" binding:"required"`
}

type FlightDTO struct {
	FlightNumber  string `json:"flight_number"`
	Airline       string `json:"airline"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	DepartureTime string `json:"departure_time"`
	ArrivalTime   string `json:"arrival_time"`
}

type FlightResponse struct {
	ID        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DelatedAt *time.Time
	FlightDTO
}
