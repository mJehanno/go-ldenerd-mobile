package models

import "fmt"

type Coins struct {
	Copper   int `json:"Copper,omitempty"`
	Silver   int `json:"Silver,omitempty"`
	Electrum int `json:"Electrum,omitempty"`
	Gold     int `json:"Gold,omitempty"`
	Platinum int `json:"Platinum,omitempty"`
}

func (c Coins) String() string {
	return fmt.Sprintf(
		"%v Platinum coins, %v Gold coins, %v Electrum coins, %v Silver coins, %v Copper coins",
		c.Platinum,
		c.Gold,
		c.Electrum,
		c.Silver,
		c.Copper,
	)
}
