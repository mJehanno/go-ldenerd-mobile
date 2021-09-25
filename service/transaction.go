package service

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

const (
	Copper Currency = iota
	Silver
	Electrum
	Gold
	Platinum
	Limit
)
