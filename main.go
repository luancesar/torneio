package main

import (
	"torneio/database"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	database.InitDB()
	// tourney.NewTourney("Torneio de Futebol")
	// tourney.CreateTolken(1, "token1")
}
