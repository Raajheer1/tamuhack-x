package flight

import "C"
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
	FlightNumber        uint      `json:"flight_number"`
	Origin              string    `json:"origin"`
	Dest                string    `json:"dest"`
	AircraftType        string    `json:"aircraft_type"`
	ScheduledFlightTime time.Time `json:"scheduled_flight_time"`
}

func CreateFlight(w http.ResponseWriter, r *http.Request) {
	// Read body into CreateRequest struct
	var request UpsertRequest
	req := json.NewDecoder(r.Body)
	err := req.Decode(&request)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid request body", w)
		return
	}

	flight := &models.Flight{
		ID:                  request.FlightNumber,
		Origin:              request.Origin,
		Dest:                request.Dest,
		AircraftType:        request.AircraftType,
		ScheduledFlightTime: request.ScheduledFlightTime,
	}

	// Create flight in database
	err = flight.Create(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error creating flight", w)
		return
	}

	utils.Respond(http.StatusCreated, "Flight created successfully", w)
}

func GetAllFlights(w http.ResponseWriter, r *http.Request) {
	flights, err := models.GetAllFlights(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting flights", w)
		return
	}

	// Convert flights to string
	flightsString, err := json.Marshal(flights)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error marshalling flights", w)
		return
	}

	utils.Respond(http.StatusOK, string(flightsString), w)
}

func GetFlight(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid flight id", w)
		return
	}

	flight := &models.Flight{
		ID: uint(idInt),
	}
	err = flight.Get(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting flight", w)
		return
	}

	// Convert flight to string
	flightString, err := json.Marshal(flight)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error marshalling flight", w)
		return
	}

	utils.Respond(http.StatusOK, string(flightString), w)
}

func UpdateFlight(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid flight id", w)
		return
	}

	flight := &models.Flight{
		ID: uint(idInt),
	}
	err = flight.Get(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error getting flight", w)
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

	flight.ID = request.FlightNumber
	flight.Origin = request.Origin
	flight.Dest = request.Dest
	flight.AircraftType = request.AircraftType
	flight.ScheduledFlightTime = request.ScheduledFlightTime

	// Update flight in database
	err = flight.UpdateFlight(database.DB)
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error updating flight", w)
		return
	}

	utils.Respond(http.StatusOK, "Flight updated successfully", w)
}

func DeleteFlight(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.Respond(http.StatusBadRequest, "Invalid flight id", w)
		return
	}

	err = models.DeleteFlight(database.DB, uint(idInt))
	if err != nil {
		utils.Respond(http.StatusInternalServerError, "Error deleting flight", w)
		return
	}

	utils.Respond(http.StatusOK, "Flight deleted successfully", w)
}
