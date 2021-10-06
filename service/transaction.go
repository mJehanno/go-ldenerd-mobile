package service

import (
	"reflect"

	"github.com/mjehanno/go-ldenerd-mobile/models"
)

type Transaction struct {
	Id     string `json:"_id,omitempty"`
	Type   TransactionType
	Amount []Coin
	IsGem  bool
	Reason string
}

type TransactionType int

const (
	Debit TransactionType = iota
	Credit
)

func (t TransactionType) String() string {
	switch t {
	case Debit:
		return "Debit"
	case Credit:
		return "Credit"
	}
	return "unknown"
}

func StringToTransactionType(s string) TransactionType {
	switch s {
	case "Debit":
		return Debit
	case "Credit":
		return Credit
	}
	return 0
}

type Coin struct {
	Value    int `json:"Value,omitempty"`
	Currency Currency
}

type Currency int

func (c Currency) String() string {
	r := ""
	switch c {
	case Copper:
		r = "Copper"
	case Silver:
		r = "Silver"
	case Electrum:
		r = "Electrum"
	case Gold:
		r = "Gold"
	case Platinum:
		r = "Platinum"
	}
	return r
}

const (
	Copper Currency = iota
	Silver
	Electrum
	Gold
	Platinum
	Limit
)

func ConvertSumOfAmountToCoin(amounts []Coin) models.Coins {
	var c models.Coins

	reflectValue := reflect.ValueOf(&c)
	coinValue := reflectValue.Elem()
	for _, a := range amounts {
		currentCoin := coinValue.FieldByName(a.Currency.String())
		sum := currentCoin.Int() + int64(a.Value)
		currentCoin.SetInt(sum)
	}
	return c
}
