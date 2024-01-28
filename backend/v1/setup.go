package v1

import (
	"github.com/Raajheer1/tamuhack-x/m/v2/v1/flight"
	"github.com/Raajheer1/tamuhack-x/m/v2/v1/seat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewServer() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", handleMain)

	r.Route("/flight", flight.Router)
	r.Route("/seat", seat.Router)

	return r
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
