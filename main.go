package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/e-commerce-backend/app/database"
	"github.com/e-commerce-backend/app/database/query"
	"github.com/e-commerce-backend/app/model"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func main() {

	database.InitDBConnection()

	var DB = database.DBConnection()

	// check ping
	fmt.Println("Check ping in MAIN")

	pingErr := DB.Ping()

	if pingErr != nil {
		log.Fatal("Ping in main error", pingErr)
	}

	fmt.Println("Connect successfully!")

	// checking the save new user
	newUser := model.User{
		FirstName: "Harry",
		LastName:  "Potter",
		Email:     "harry@yahoo.com",
		Phone:     "111 111-2222",
		Address:   "111 Horward Street",
	}

	query.SaveNewUser(newUser)

	// check the get user

	resultUsers, err := query.GetUserByID(10)

	if err != nil {
		log.Fatal("Error query user by id ", err)
	}

	fmt.Println(resultUsers)
	defer DB.Close()
}
