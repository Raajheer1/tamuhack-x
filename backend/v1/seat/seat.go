package seat

import (
	"encoding/json"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database/models"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

type UpsertRequest struct {
	FlightID    uint      `json:"flight_id"`
	SeatNo      string    `json:"seat_no"`
	AssignedPNR string    `json:"assigned_pnr"`
	ForSale     bool      `json:"for_sale"`
	MinPrice    uint      `json:"min_price"` // All prices in Dollars
	BuyNowPrice uint      `json:"buy_now_price"`
	CurrentBid  uint      `json:"current_bid"`
	LastBidder  string    `json:"last_bidder"`
	BidTimeEnd  time.Time `json:"bid_time_end"`
}

func CreateSeat(w http.ResponseWriter, r *http.Request) {
	// Read body into CreateRequest struct
	var request UpsertRequest
	req := json.NewDecoder(r.Body)
	err := req.Decode(&request)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid request body", w)
		return
	}

	seat := &models.Seat{
		FlightID:    request.FlightID,
		SeatNo:      request.SeatNo,
		AssignedPNR: request.AssignedPNR,
		ForSale:     request.ForSale,
		MinPrice:    request.MinPrice,
		BuyNowPrice: request.BuyNowPrice,
		CurrentBid:  request.CurrentBid,
		LastBidder:  request.LastBidder,
		BidTimeEnd:  request.BidTimeEnd,
	}

	// Create seat in database
	err = seat.Create(database.DB, seat)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error creating seat", w)
		return
	}

	utils.Respond(http.StatusCreated, "Seat created successfully", w)
}

func GetAllSeats(w http.ResponseWriter, r *http.Request) {
	seats, err := models.GetAllSeats(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting seats", w)
		return
	}

	// Convert seats to string
	seatsString, err := json.Marshal(seats)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error marshalling seats", w)
		return
	}

	utils.Respond(http.StatusOK, string(seatsString), w)
}

func GetAllSeatsWithFlightId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid flight id", w)
		return
	}

	seats, err := models.GetAllSeatsWithFlightID(database.DB, uint(idInt))
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting seats", w)
		return
	}

	// Convert seats to string
	seatsString, err := json.Marshal(seats)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error marshalling seats", w)
		return
	}

	utils.Respond(http.StatusOK, string(seatsString), w)
}

func GetSeat(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid seat id", w)
		return
	}

	seat, err := models.GetSeat(database.DB, uint(idInt))
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting seat", w)
		return
	}

	// Convert seat to string
	seatString, err := json.Marshal(seat)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error marshalling seat", w)
		return
	}

	utils.Respond(http.StatusOK, string(seatString), w)
}

func UpdateSeat(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid seat id", w)
		return
	}

	seat, err := models.GetSeat(database.DB, uint(idInt))
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting seat", w)
		return
	}

	// Read body into CreateRequest struct
	var request UpsertRequest
	req := json.NewDecoder(r.Body)
	err = req.Decode(&request)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid request body", w)
		return
	}

	seat.FlightID = request.FlightID
	seat.SeatNo = request.SeatNo
	seat.AssignedPNR = request.AssignedPNR
	seat.ForSale = request.ForSale
	seat.MinPrice = request.MinPrice
	seat.BuyNowPrice = request.BuyNowPrice
	seat.CurrentBid = request.CurrentBid
	seat.LastBidder = request.LastBidder
	seat.BidTimeEnd = request.BidTimeEnd

	// Update seat in database
	err = seat.Update(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error updating seat", w)
		return
	}

	utils.Respond(http.StatusOK, "Seat updated successfully", w)
}

type SwapRequest struct {
	SeatID1 uint `json:"seat_id_1"`
	SeatID2 uint `json:"seat_id_2"`
}

func SwapSeats(w http.ResponseWriter, r *http.Request) {
	var request SwapRequest
	req := json.NewDecoder(r.Body)
	err := req.Decode(&request)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid request body", w)
		return
	}

	seat1, err := models.GetSeat(database.DB, request.SeatID1)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting seat 1", w)
		return
	}

	seat2, err := models.GetSeat(database.DB, request.SeatID2)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting seat 2", w)
		return
	}

	seat1.AssignedPNR, seat2.AssignedPNR = seat2.AssignedPNR, seat1.AssignedPNR
	seat1.ForSale, seat2.ForSale = false, false
	seat1.MinPrice, seat2.MinPrice = 0, 0
	seat1.BuyNowPrice, seat2.BuyNowPrice = 0, 0
	seat1.CurrentBid, seat2.CurrentBid = 0, 0
	seat1.LastBidder, seat2.LastBidder = "", ""

	err = seat1.Update(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error updating seat 1", w)
		return
	}

	err = seat2.Update(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error updating seat 2", w)
		return
	}

	utils.Respond(http.StatusOK, "Seats swapped successfully", w)
}

func DeleteSeat(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid seat id", w)
		return
	}

	err = models.DeleteSeat(database.DB, uint(idInt))
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error deleting seat", w)
		return
	}

	utils.Respond(http.StatusOK, "Seat deleted successfully", w)
}
