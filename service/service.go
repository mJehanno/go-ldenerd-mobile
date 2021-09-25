package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"fyne.io/fyne/v2/data/binding"
	"github.com/mjehanno/go-ldenerd-mobile/auth"
)

var BaseUrl binding.String = binding.NewString()
var token string
var Auth auth.Auth

var tokenBinding = binding.NewString()

func Login(username, password string) error {

	values := map[string]string{"Username": username, "Password": password}
	json_data, err := json.Marshal(values)

	url, _ := BaseUrl.Get()
	resp, err := http.Post(url+"/api/login", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Status Code := %v \n", resp.StatusCode)
	if resp.StatusCode != 200 {
		return errors.New("failed to login")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &Auth)
	fmt.Println(Auth)
	tokenBinding.Set(Auth.AccessToken)

	go RefreshToken()
	defer resp.Body.Close()

	return nil
}

func RefreshToken() {

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
	fmt.Printf("Get Gold - Status : %v \n", resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println(string(body))
		return string(body)
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
	fmt.Printf("Get Gold Detail - Status : %v \n", resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println(string(body))
		return string(body)
	}

	return ""
}

func GetHistory() {
	//url, _ := BaseUrl.Get()
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
