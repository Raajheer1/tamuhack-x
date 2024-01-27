package models

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
