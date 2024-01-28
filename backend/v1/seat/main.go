package seat

import "github.com/go-chi/chi/v5"

func Router(r chi.Router) {
	r.Post("/", CreateSeat)
	r.Get("/", GetAllSeats)
	r.Get("/flight/{id}", GetAllSeatsWithFlightId)
	r.Post("/swap", SwapSeats)
	r.Put("/{id}", UpdateSeat)
	r.Delete("/{id}", DeleteSeat)
}
