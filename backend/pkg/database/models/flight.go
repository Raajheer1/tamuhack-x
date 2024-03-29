package models

import (
	"gorm.io/gorm"
	"time"
)

type Flight struct {
	ID                  uint      `json:"id" gorm:"primaryKey"`
	Origin              string    `json:"origin"`
	Dest                string    `json:"destination"`
	AircraftType        string    `json:"aircraft_type"`
	ScheduledDepartureTime time.Time `json:"scheduled_departure_time"`
	ScheduledArrivalTime time.Time `json:"scheduled_arrival_time"`
}

func (f *Flight) Create(db *gorm.DB) error {
	err := db.Create(f).Error
	if err != nil {
		return err
	}

	return nil
}


func (f *Flight) Get(db *gorm.DB) error {
	return db.Where("id = ?", f.ID).First(f).Error
}

func GetAllFlights(db *gorm.DB) ([]*Flight, error) {
	var flights []*Flight
	err := db.Find(&flights).Error
	if err != nil {
		return nil, err
	}

	return flights, nil
}

func (f *Flight) UpdateFlight(db *gorm.DB) error {
	err := db.Save(f).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteFlight(db *gorm.DB, id uint) error {
	err := db.Delete(&Flight{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
