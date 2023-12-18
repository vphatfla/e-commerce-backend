package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/e-commerce-backend/app/database"
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

}
