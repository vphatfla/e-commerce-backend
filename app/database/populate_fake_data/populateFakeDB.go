package populatefakedata

import (
	"log"

	"github.com/e-commerce-backend/app/auth"
	"github.com/e-commerce-backend/app/database/query"
	"github.com/e-commerce-backend/app/model"
)

func PopulateDummyUsers() {
	// checking the save new user
	tempHashed, _ := auth.HashPassword("password123")
	// Sample user records
	var users = []model.User{
		{
			FirstName: "Harry",
			LastName:  "Potter",
			Email:     "harry@yahoo.com",
			Phone:     "111 111-2222",
			Address:   "111 Horward Street",
			Password:  tempHashed,
		},
		{
			FirstName: "Hermione",
			LastName:  "Granger",
			Email:     "hermione@gmail.com",
			Phone:     "222 222-3333",
			Address:   "222 Granger Lane",
			Password:  tempHashed,
		},
		{
			FirstName: "Ron",
			LastName:  "Weasley",
			Email:     "ron@hotmail.com",
			Phone:     "333 333-4444",
			Address:   "333 Weasley Road",
			Password:  tempHashed,
		},
		{
			FirstName: "Ginny",
			LastName:  "Weasley",
			Email:     "ginny@gmail.com",
			Phone:     "444 444-5555",
			Address:   "444 Weasley Avenue",
			Password:  tempHashed,
		},
		{
			FirstName: "Albus",
			LastName:  "Dumbledore",
			Email:     "albus@yahoo.com",
			Phone:     "555 555-6666",
			Address:   "555 Dumbledore Tower",
			Password:  tempHashed,
		},
	}

	for _, user := range users {
		err := query.SaveNewUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func PopulateDummyCart() {

}
