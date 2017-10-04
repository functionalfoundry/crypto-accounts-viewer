package data

import (
	"time"
)

// A crypto-currency exchange (e.g. GDAX).
type Exchange struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A currency with a name (e.g. USD, ETH, BTC).
type Currency struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LongName string `json:"longName"`
}

// A transfer in an account.
type Transfer struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Time      time.Time `json:"time"`
	Amount    float64   `json:"amount"`
	Balance   float64   `json:"balance"`
	Account   Account   `json:"account"`
	AccountID int       `json:"accountId"`
}

// A crypto-currency account on one of the known exchanges.
type Account struct {
	ID         int         `json:"id"`
	Currency   Currency    `json:"currency"`
	CurrencyID int         `json:"currencyId"`
	Exchange   Exchange    `json:"exchange"`
	ExchangeID int         `json:"exchangeId"`
	Transfers  []*Transfer `json:"transferIds"`
}
