package database

import (
	"github.com/go-pg/pg"

	"../data"
)

func CreateAccount(db *pg.DB, args map[string]interface{}) (interface{}, error) {
	account := data.Account{
		CurrencyID: args["currencyId"].(int),
		ExchangeID: args["exchangeId"].(int),
	}
	err := db.Insert(&account)
	return &account, err
}
