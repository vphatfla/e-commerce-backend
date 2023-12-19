package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type hashConfig struct {
	Cost int `json:"cost"`
}

var bcrypt_cost int

func InitHashing() {
	jsonFile, err := os.Open("config/bcrypt_config.json")

	if err != nil {
		log.Fatal("Error opening bcrypt config  ", err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal("Error converting []byte", err)
	}

	var hashCfg hashConfig

	err = json.Unmarshal([]byte(byteValue), &hashCfg)

	if err != nil {
		log.Fatal("Error unmarshal json")
	}

	bcrypt_cost = hashCfg.Cost
}

func HashPassword(rawPassword string) ([]byte, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt_cost)

	if err != nil {
		return nil, err
	}
	return hashedByte, nil
}
