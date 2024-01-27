package models

type Flight struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Origin       string `json:"origin"`
	Dest         string `json:"dest"`
	AircraftType string `json:"aircraft_type"`
}
