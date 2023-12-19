package main

import (
	"database/sql"

	"github.com/e-commerce-backend/app/auth"
	"github.com/e-commerce-backend/app/database"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func main() {

	database.InitDBConnection()
	auth.InitHashing()
	// var DB = database.DBConnection()

	// // check ping
	// fmt.Println("Check ping in MAIN")

	// pingErr := DB.Ping()

	// if pingErr != nil {
	// 	log.Fatal("Ping in main error", pingErr)
	// }

	// fmt.Println("Connect successfully!")

	// checking the save new user
	// tempHasheed, _ := auth.HashPassword("password123")
	// newUser := model.User{
	// 	FirstName: "Harry",
	// 	LastName:  "Potter",
	// 	Email:     "harry@yahoo.com",
	// 	Phone:     "111 111-2222",
	// 	Address:   "111 Horward Street",
	// 	Password:  tempHasheed,
	// }

	// err := query.SaveNewUser(newUser)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// check the get user

	// resultUsers, err := query.GetUserByID(10)

	// if err != nil {
	// 	log.Fatal("Error query user by id ", err)
	// }

	// fmt.Println(resultUsers)

	// // check the get user by first name

	// resultUsers1, err := query.GetUsersByFirstName("Harry")

	// for _, u := range resultUsers1 {
	// 	fmt.Println(u)
	// }

	// saveErr := query.SaveUserPassword("example@gmail.com", "password123")

	// if saveErr != nil {
	// 	log.Fatal("Error saving user password ", saveErr)
	// }

	//defer DB.Close()
}
