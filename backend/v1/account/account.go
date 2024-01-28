package account

import (
	"encoding/json"
	"net/http"

	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database/models"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/utils"
	"github.com/go-chi/chi/v5"
)

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

	for i := range seats {
		seats[i].Account = account

		f := models.Flight{
			ID: seats[i].FlightID,
		}

		err = f.Get(database.DB)
		if err != nil {
			utils.Respond(http.StatusInternalServerError, "Error getting flight", w)
			return
		}

		seats[i].Flight = f
	}

	seatString, err := json.Marshal(seats)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error marshalling seats", w)
		return
	}

	utils.Respond(http.StatusOK, string(seatString), w)
}

type BidRequest struct {
	BoughtSeatID uint `json:"bought_seat_id"`
	TradedSeatID uint `json:"traded_seat_id"`
	Bid          uint `json:"bid"`
}

func Bid(w http.ResponseWriter, r *http.Request) {
	var request BidRequest
	req := json.NewDecoder(r.Body)
	err := req.Decode(&request)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid request body", w)
		return
	}

	boughtSeat, err := models.GetSeat(database.DB, request.BoughtSeatID)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting bought seat", w)
		return
	}

	tradedSeat, err := models.GetSeat(database.DB, request.TradedSeatID)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting traded seat", w)
		return
	}

	if boughtSeat.CurrentBid >= request.Bid {
		utils.Respond(http.StatusBadRequest, "Bid must be greater than current bid", w)
		return
	}

	buyer := models.Account{
		ID: tradedSeat.AccountID,
	}

	err = buyer.Get(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting account", w)
		return
	}

	seller := models.Account{
		ID: boughtSeat.AccountID,
	}

	err = seller.Get(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting account", w)
		return
	}

	if boughtSeat.CurrentBid >= boughtSeat.BuyNowPrice {
		buyer.Money -= boughtSeat.CurrentBid
		seller.Money += boughtSeat.CurrentBid
		if err := buyer.UpdateAccount(database.DB); err != nil {
			utils.Respond(http.StatusInternalServerError, "Error updating buyer account", w)
			return
		}

		if err := seller.UpdateAccount(database.DB); err != nil {
			utils.Respond(http.StatusInternalServerError, "Error updating seller account", w)
			return
		}

		boughtSeat.AccountID, tradedSeat.AccountID = tradedSeat.AccountID, boughtSeat.AccountID
		boughtSeat.CurrentBid, tradedSeat.CurrentBid = 0, 0
		boughtSeat.LastBidder, tradedSeat.LastBidder = "", ""
		boughtSeat.ForSale, tradedSeat.ForSale = false, false
	} else {
		boughtSeat.CurrentBid = request.Bid
		boughtSeat.LastBidder = tradedSeat.AccountID
	}

	err = boughtSeat.Update(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error updating bought seat", w)
		return
	}

	err = tradedSeat.Update(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error updating traded seat", w)
		return
	}

	utils.Respond(http.StatusOK, "Bid successful", w)
}

type BuyRequest struct {
	BoughtSeatID uint `json:"bought_seat_id"`
	TradedSeatID uint `json:"traded_seat_id"`
}

func Buy(w http.ResponseWriter, r *http.Request) {
	var request BuyRequest
	req := json.NewDecoder(r.Body)
	err := req.Decode(&request)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid request body", w)
		return
	}

	boughtSeat, err := models.GetSeat(database.DB, request.BoughtSeatID)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting bought seat", w)
		return
	}

	tradedSeat, err := models.GetSeat(database.DB, request.TradedSeatID)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting traded seat", w)
		return
	}

	buyer := models.Account{
		ID: tradedSeat.AccountID,
	}

	err = buyer.Get(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting account", w)
		return
	}

	seller := models.Account{
		ID: boughtSeat.AccountID,
	}

	err = seller.Get(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting account", w)
		return
	}

	buyer.Money -= boughtSeat.BuyNowPrice
	seller.Money += boughtSeat.BuyNowPrice
	if err := buyer.UpdateAccount(database.DB); err != nil {
		utils.Respond(http.StatusInternalServerError, "Error updating buyer account", w)
		return
	}

	if err := seller.UpdateAccount(database.DB); err != nil {
		utils.Respond(http.StatusInternalServerError, "Error updating seller account", w)
		return
	}

	boughtSeat.AccountID, tradedSeat.AccountID = tradedSeat.AccountID, boughtSeat.AccountID
	boughtSeat.CurrentBid, tradedSeat.CurrentBid = 0, 0
	boughtSeat.LastBidder, tradedSeat.LastBidder = "", ""
	boughtSeat.ForSale, tradedSeat.ForSale = false, false

	err = boughtSeat.Update(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error updating bought seat", w)
		return
	}

	err = tradedSeat.Update(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error updating traded seat", w)
		return
	}

	utils.Respond(http.StatusOK, "Buy now successful", w)
}
