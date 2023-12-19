package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hostname string `json:"hostname"`
	Dbname   string `json:"dbname"`
	Net      string `json:"net"`
}

var DB *sql.DB

func InitDBConnection() {
	jsonFile, err := os.Open("config/mysql_config.json")

	if err != nil {
		log.Fatal("Error open config file", err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal("Error converting []byte", err)
	}

	var dbConfig DBConfig

	err = json.Unmarshal([]byte(byteValue), &dbConfig)

	if err != nil {
		log.Fatal("Error unmarshal json")
	}

	cfgFormatted := mysql.Config{
		User:   dbConfig.Username,
		Passwd: dbConfig.Password,
		Net:    dbConfig.Net,
		Addr:   dbConfig.Hostname,
		DBName: dbConfig.Dbname,
	}

	DB, err = sql.Open("mysql", cfgFormatted.FormatDSN())

	if err != nil {
		log.Fatal("Error connect to MYSQL SERVER", err)
	}

	pingErr := DB.Ping()

	if pingErr != nil {
		log.Fatal("Error pinging MYSQL SERVER", pingErr)
	}

	fmt.Println("Connected to MYSQL SERVER successfully!!!")

	defer jsonFile.Close()
}
func DBConnection() *sql.DB {
	return DB
}
