package router

import (
	"torneio/pkg/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/tourneys/{id}", handlers.GetTourney).Methods("GET")
	r.HandleFunc("/api/tourneys/{id}", handlers.CreateTourneyTolken).Methods("POST")
	r.HandleFunc("/api/tourneys", handlers.CreateTourney).Methods("POST")

	return r
}
