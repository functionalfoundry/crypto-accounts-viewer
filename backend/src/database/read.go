package database

import (
	"github.com/go-pg/pg"

	"../data"
)

func GetAccounts(db *pg.DB) ([]*data.Account, error) {
	var accounts []*data.Account
	err := db.Model(&accounts).
		Column("account.*", "Currency", "Exchange").
		Select()
	return accounts, err
}

func GetCurrencies(db *pg.DB) ([]*data.Currency, error) {
	var currencies []*data.Currency
	err := db.Model(&currencies).
		Column("currency.*").
		Select()
	return currencies, err
}

func GetExchanges(db *pg.DB) ([]*data.Exchange, error) {
	var exchanges []*data.Exchange
	err := db.Model(&exchanges).
		Column("exchange.*").
		Select()
	return exchanges, err
}
