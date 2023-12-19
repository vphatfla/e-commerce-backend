package query

import (
	"github.com/e-commerce-backend/app/database"
	"github.com/e-commerce-backend/app/model"
)

func GetUserByID(id int) ([]model.User, error) {
	db := database.DBConnection()
	query := "SELECT * FROM User WHERE UserID = ?;"

	rows, err := db.Query(query, id)

	if err != nil {
		return nil, err
	}

	// loops through rows
	var users []model.User

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Address, &user.Phone); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func GetUsersByFirstName(firstName string) ([]model.User, error) {
	db := database.DBConnection()
	query := "SELECT * FROM User WHERE FirstName = ?;"

	rows, err := db.Query(query, firstName)

	if err != nil {
		return nil, err
	}

	// loops through rows
	var users []model.User

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Address, &user.Phone); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func GetUsersByLastName(lastName string) ([]model.User, error) {
	db := database.DBConnection()
	query := "SELECT * FROM User WHERE FirstName = ?;"

	rows, err := db.Query(query, lastName)

	if err != nil {
		return nil, err
	}

	// loops through rows
	var users []model.User

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Address, &user.Phone); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func GetUserByEmail(email string) ([]model.User, error) {
	db := database.DBConnection()
	query := "SELECT * FROM User WHERE FirstName = ?;"

	rows, err := db.Query(query, email)

	if err != nil {
		return nil, err
	}

	// loops through rows
	var users []model.User

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Address, &user.Phone); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func GetUsersByAddress(address string) ([]model.User, error) {
	db := database.DBConnection()
	query := "SELECT * FROM User WHERE FirstName = ?;"

	rows, err := db.Query(query, address)

	if err != nil {
		return nil, err
	}

	// loops through rows
	var users []model.User

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Address, &user.Phone); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}
