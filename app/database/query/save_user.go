package query

import (
	"fmt"
	"log"

	"github.com/e-commerce-backend/app/database"
	"github.com/e-commerce-backend/app/model"
)

func SaveNewUser(user model.User) {
	db := database.DBConnection()

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal("Ping error inside save_user  ", pingErr)
	}

	fmt.Println("Save_user: connection OK")

	// build query
	query := "INSERT INTO User (FirstName, LastName, Email, Address, Phone) VALUES (?,?,?,?,?);"

	insertResult, err := db.Exec(query, user.FirstName, user.LastName, user.Email, user.Address, user.Phone)

	if err != nil {
		log.Fatal("Error inserting ", user.FirstName, " ---  ", err)
	}

	fmt.Println("Success!", insertResult)
}
