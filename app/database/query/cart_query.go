package query

import (
	"github.com/e-commerce-backend/app/database"
	"github.com/e-commerce-backend/app/model"
)

func SaveNewCart(newCart model.Cart) error {
	db := database.DBConnection()

	pingErr := db.Ping()

	if pingErr != nil {
		return pingErr
	}

	query := "INSERT INTO Cart (UserID, Total) VALUES (?,?);"

	_, err := db.Exec(query, newCart.UserID, newCart.Total)

	if err != nil {
		return err
	}

	return nil
}

func DeleteCart() {

}

func UpdateCart() {

}
