package flight

import "github.com/go-chi/chi/v5"

func Router(r chi.Router) {
	r.Post("/", CreateFlight)
	r.Get("/", GetAllFlights)
	r.Get("/{id}", GetFlight)
	r.Put("/{id}", UpdateFlight)
	r.Delete("/{id}", DeleteFlight)
}
