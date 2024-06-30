package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"torneio/tourney"

	"github.com/gorilla/mux"
)

func CreateTourneyTolken(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var body struct{ TolkenName string }

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	tourney.CreateTolken(id, body.TolkenName)
}

func GetTourney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	tourney := tourney.GetTourney(id)
	var blockchain map[string]interface{}
	json.Unmarshal([]byte(tourney.Blockchain), &blockchain)

	response := struct {
		Name       string                 `json:"Name"`
		Blockchain map[string]interface{} `json:"Blockchain"`
	}{
		Name:       tourney.Name,
		Blockchain: blockchain,
	}

	json.NewEncoder(w).Encode(response)
}

func CreateTourney(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var body struct{ Name string }

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tourney.NewTourney(body.Name)
}
