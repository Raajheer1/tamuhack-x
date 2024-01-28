package seat

import "github.com/go-chi/chi/v5"

func Router(r chi.Router) {
	r.Post("/", CreateSeat)
	r.Get("/", GetAllSeats)
	r.Get("/{id}", GetSeat)
	r.Post("/swap", SwapSeats)
	r.Put("/{id}", UpdateSeat)
	r.Delete("/{id}", DeleteSeat)
}
