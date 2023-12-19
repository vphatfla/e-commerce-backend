package query

import (
	"fmt"

	"github.com/e-commerce-backend/app/database"
	"github.com/e-commerce-backend/app/model"
)

func SaveNewUser(user model.User) error {
	db := database.DBConnection()

	pingErr := db.Ping()

	if pingErr != nil {
		return pingErr
	}

	// build query
	query := "INSERT INTO User (FirstName, LastName, Email, Address, Phone, HashedPassword) VALUES (?,?,?,?,?,?);"

	insertResult, err := db.Exec(query, user.FirstName, user.LastName, user.Email, user.Address, user.Phone, user.Password)

	if err != nil {
		return err
	}

	fmt.Println("Success!", insertResult)

	return nil
}
