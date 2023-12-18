package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hostname string `json:"hostname"`
	Dbname   string `json:"dbname"`
	Net      string `json:"net"`
}

var DB *sql.DB

func main() {
	// database file
	jsonFile, err := os.Open("config/mysql_config.json")

	if err != nil {
		fmt.Println(err)
	}

	// byte[]
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}

	// parse json
	var dbConfig DBConfig

	err = json.Unmarshal([]byte(byteValue), &dbConfig)

	if err != nil {
		fmt.Println(err)
	}

	// mysql config

	cfgFormatted := mysql.Config{
		User:   dbConfig.Username,
		Passwd: dbConfig.Password,
		Net:    dbConfig.Net,
		Addr:   dbConfig.Hostname,
		DBName: dbConfig.Dbname,
	}

	DB, err = sql.Open("mysql", cfgFormatted.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to MYSQL server successfully!")
	defer jsonFile.Close()
}
