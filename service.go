package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var BaseUrl string
var token string

func Login(username, password string) error {

	values := map[string]string{"login": username, "password": password}
	json_data, err := json.Marshal(values)

	resp, err := http.Post(BaseUrl+"/login", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}

func RefreshToken() {

}

func GetGold() {

}

func GetGoldDetail() {

}

func GetHistory() {

}

func AddTransaction() {

}
