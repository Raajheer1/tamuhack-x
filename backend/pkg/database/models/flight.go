package models

import "gorm.io/gorm"

type Flight struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Origin       string `json:"origin"`
	Dest         string `json:"dest"`
	AircraftType string `json:"aircraft_type"`
}

func (f *Flight) Create(db *gorm.DB, flight *Flight) error {
	err := db.Create(flight).Error
	if err != nil {
		return err
	}

	return nil
}

func (f *Flight) Get(db *gorm.DB, id uint) (*Flight, error) {
	flight := &Flight{}
	err := db.First(flight, id).Error
	if err != nil {
		return nil, err
	}

	return flight, nil
}

func (f *Flight) GetAll(db *gorm.DB) ([]*Flight, error) {
	var flights []*Flight
	err := db.Find(&flights).Error
	if err != nil {
		return nil, err
	}

	return flights, nil
}

func (f *Flight) Update(db *gorm.DB, flight *Flight) error {
	err := db.Save(flight).Error
	if err != nil {
		return err
	}

	return nil
}

func (f *Flight) Delete(db *gorm.DB, id uint) error {
	err := db.Delete(&Flight{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
