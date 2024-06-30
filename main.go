package main

import (
	"log"
	"net/http"
	"torneio/database"
	"torneio/pkg/router"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	database.InitDB()

	r := router.NewRouter()

	// Iniciar o servidor
	log.Fatal(http.ListenAndServe(":8000", r))
}
