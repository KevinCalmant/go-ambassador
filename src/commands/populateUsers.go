package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/bxcodec/faker/v3"
)

// If using docker we need to connect to the container sh to run the command.
func main() {
	database.Connect()
	for i := 0; i < 30; i++ {
		ambassador := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}
		ambassador.SetPassword("1234")
		database.DB.Create(&ambassador)
	}
}