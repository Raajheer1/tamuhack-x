package models

import (
	"gorm.io/gorm"
	"time"
)

type Seat struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	FlightID    uint      `json:"flight_id"`
	Flight      Flight    `json:"flight" gorm:"foreignKey:FlightID"`
	SeatNo      string    `json:"seat_number"`
	AccountID   string    `json:"account_id" gorm:"type:varchar(10)"`
	Account     Account   `json:"account" gorm:"foreignKey:AccountID"`
	AssignedPNR string    `json:"assigned_pnr"`
	ForSale     bool      `json:"for_sale"`
	MinPrice    uint      `json:"min_price"` // All prices in Dollars
	BuyNowPrice uint      `json:"buy_now_price"`
	CurrentBid  uint      `json:"current_bid"`
	LastBidder  string    `json:"last_bidder"`
	BidTimeEnd  time.Time `json:"bid_time_end"`
}

func (s *Seat) Create(db *gorm.DB, seat *Seat) error {
	err := db.Create(seat).Error
	if err != nil {
		return err
	}

	return nil
}

func GetSeat(db *gorm.DB, id uint) (*Seat, error) {
	seat := &Seat{}
	err := db.First(seat, id).Error
	if err != nil {
		return nil, err
	}

	return seat, nil
}

func GetAllSeats(db *gorm.DB) ([]*Seat, error) {
	var seats []*Seat
	err := db.Find(&seats).Error
	if err != nil {
		return nil, err
	}

	return seats, err
}

func GetAllSeatsWithAccountID(db *gorm.DB, accountID string) ([]*Seat, error) {
	var seats []*Seat
	err := db.Where("account_id = ?", accountID).Find(&seats).Error
	if err != nil {
		return nil, err
	}

	return seats, err
}

func GetAllSeatsWithFlightID(db *gorm.DB, flightID uint) ([]*Seat, error) {
	var seats []*Seat
	err := db.Where("flight_id = ?", flightID).Find(&seats).Error
	if err != nil {
		return nil, err
	}

	return seats, err
}

func (s *Seat) Update(db *gorm.DB) error {
	return db.Save(s).Error
}

func DeleteSeat(db *gorm.DB, id uint) error {
	err := db.Delete(&Seat{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
