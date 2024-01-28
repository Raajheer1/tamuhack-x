package account

import "github.com/go-chi/chi/v5"

func Router(r chi.Router) {
	r.Get("/fetchall/{id}", DoIt)
	r.Post("/bid", Bid)
	r.Post("/buy", Buy)
}
