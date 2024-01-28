package account

import (
	"encoding/json"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database/models"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Response struct {
	AAvantage string `json:"aadvantage_number"`
}

func DoIt(w http.ResponseWriter, r *http.Request) {
	AAid := chi.URLParam(r, "id")

	account := models.Account{
		ID: AAid,
	}
	err := account.Get(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting account", w)
		return
	}

	seats, err := models.GetAllSeatsWithAccountID(database.DB, account.ID)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting seats", w)
		return
	}

	seatString, err := json.Marshal(seats)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error marshalling seats", w)
		return
	}

	utils.Respond(http.StatusOK, string(seatString), w)
}

type BidRequest struct {
	SeatID   uint   `json:"seat_id"`
	BidderID string `json:"bidder_id"`
	Bid      uint   `json:"bid"`
}

func Bid(w http.ResponseWriter, r *http.Request) {
	var request BidRequest
	req := json.NewDecoder(r.Body)
	err := req.Decode(&request)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid request body", w)
		return
	}

	seat, err := models.GetSeat(database.DB, request.SeatID)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting seat", w)
		return
	}

	if seat.CurrentBid >= request.Bid {
		utils.Respond(http.StatusBadRequest, "Bid must be higher than current bid", w)
		return
	}

	seat.CurrentBid = request.Bid
	seat.LastBidder = request.BidderID
	err = models.UpdateSeat(database.DB, seat)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error updating seat", w)
		return
	}

	utils.Respond(http.StatusOK, "Bid successful", w)
}
