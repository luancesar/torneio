package tourney

import (
	"time"
	"torneio/blockchain"
	"torneio/database"
	"torneio/models"

	"encoding/json"
)

type Tourney struct {
	Name       string
	Blockchain string
}

func NewTourney(name string) {

	data, _ := json.MarshalIndent(blockchain.NewBlockchain(), "", "  ")
	tourney := Tourney{
		Name:       name,
		Blockchain: string(data),
	}
	database.DB.Create(&tourney)
}

func GetTourneys() []models.Tourney {

	var tourneys []models.Tourney
	database.DB.Find(&tourneys)
	return tourneys
}

func GetTourney(id int) models.Tourney {

	var tourney models.Tourney
	database.DB.First(&tourney, id)
	return tourney
}

func UpdateTourney(id int, bc string) {

	database.DB.Model(&models.Tourney{}).Where("id = ?", id).Updates(map[string]interface{}{"blockchain": bc, "updated_at": time.Now()})
}

func CreateTolken(tourneyId int, name string) {

	tourney := GetTourney(tourneyId)
	var bc blockchain.Blockchain
	json.Unmarshal([]byte(tourney.Blockchain), &bc)
	bc.AddBlock(name)
	data, _ := json.MarshalIndent(bc, "", "  ")
	UpdateTourney(tourneyId, string(data))
}
