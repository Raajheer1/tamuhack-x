package models

import "gorm.io/gorm"

type Seat struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	FlightID    uint   `json:"flight_id"`
	Flight      Flight `json:"flight" gorm:"foreignKey:FlightID"`
	SeatNo      string `json:"seat_no"`
	AssignedPNR string `json:"assigned_pnr"`
	ForSale     bool   `json:"for_sale"`
	MinPrice    uint   `json:"min_price"` // All prices in Dollars
	BuyNowPrice uint   `json:"buy_now_price"`
	CurrentBid  uint   `json:"current_bid"`
}


func (s *Seat) Create(db *gorm.DB, seat *Seat) error {
	err := db.Create(seat).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Seat) Get(db *gorm.DB, id uint) (*Seat, error) {
	seat := &Seat{}
	err := db.First(seat, id).Error
	if err != nil {
		return nil, err
	}

	return seat, nil
}

func (s *Seat) GetAll(db *gorm.DB) ([]*Seat, error) {
	var seats []*Seat
	err := db.Find(&seats).Error
	if err != nil {
		return nil, err
	}

	return seats, nil
}

func (s *Seat) Update(db *gorm.DB, seat *Seat) error {
	err := db.Save(seat).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Seat) Delete(db *gorm.DB, id uint) error {
	err := db.Delete(&Seat{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
