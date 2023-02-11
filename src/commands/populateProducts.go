package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

// If using docker we need to connect to the container sh to run the command.
func main() {
	database.Connect()
	for i := 0; i < 30; i++ {
		product := models.Product{
			Title:       faker.Sentence(),
			Description: faker.Sentence(),
			Image:       faker.URL(),
			Price:       rand.Float64(),
		}
		database.DB.Create(&product)
	}
}
