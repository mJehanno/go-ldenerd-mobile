package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"fyne.io/fyne/v2/data/binding"
	"github.com/mjehanno/go-ldenerd-mobile/auth"
	"github.com/mjehanno/go-ldenerd-mobile/models"
)

var BaseUrl binding.String = binding.NewString()
var refreshTokenBinding = binding.NewString()

var tokenBinding = binding.NewString()

func Login(username, password string) error {

	values := map[string]string{"Username": username, "Password": password}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
		return err
	}
	url, _ := BaseUrl.Get()
	resp, err := http.Post(url+"/api/login", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("failed to login")
	}
	var authO auth.Auth
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &authO)
	tokenBinding.Set(authO.AccessToken)
	refreshTokenBinding.Set(authO.RefreshToken)
	go RefreshToken(authO.ExpiresIn)
	defer resp.Body.Close()

	return nil
}

func RefreshToken(refreshTime int) {
	var authO auth.Auth
	if refreshTime <= 0 {
		time.Sleep(119 * time.Second)
	} else {
		time.Sleep(time.Duration((refreshTime - 50) * int(time.Second)))
	}

	client := &http.Client{}
	url, _ := BaseUrl.Get()
	token, _ := tokenBinding.Get()
	refreshToken, _ := refreshTokenBinding.Get()
	jsonB, _ := json.Marshal(struct {
		RefreshToken string `json:"refresh_token"`
	}{RefreshToken: refreshToken})

	req, err := http.NewRequest("POST", url+"/api/refresh", bytes.NewBuffer(jsonB))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, _ := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	json.Unmarshal(body, &authO)
	tokenBinding.Set(authO.AccessToken)
	refreshTokenBinding.Set(authO.RefreshToken)
	go RefreshToken(authO.ExpiresIn)
}

func GetGold() string {
	url, _ := BaseUrl.Get()
	resp, err := http.Get(url + "/api/gold")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		var coins models.Coins
		json.Unmarshal(body, &coins)
		return fmt.Sprintf("%v Gold Coins", coins.Gold)
	}

	return ""
}

func GetGoldDetail() string {
	url, _ := BaseUrl.Get()
	resp, err := http.Get(url + "/api/gold/details")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == 200 {
		var coins models.Coins
		json.Unmarshal(body, &coins)
		return fmt.Sprintf("You curently have %v.", coins.String())
	}

	return ""
}

func GetHistory() []Transaction {
	url, _ := BaseUrl.Get()
	url += "/api/transactions/history"

	resp, _ := http.Get(url)
	var history []Transaction
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &history)
	}
	return history
}

func AddTransaction(t Transaction) {
	url, _ := BaseUrl.Get()
	url += "/api/transactions"
	token, _ := tokenBinding.Get()
	client := &http.Client{}

	jsonBody, _ := json.Marshal(&t)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 201 {
		fmt.Printf("Error while creating transaction : Code = %v, body = %v \n", resp.StatusCode, string(body))
	}

}
