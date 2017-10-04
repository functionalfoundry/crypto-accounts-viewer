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

func GetCurrencies(db *pg.DB) ([]data.Currency, error) {
	return []data.Currency{}, nil
}

func GetExchanges(db *pg.DB) ([]data.Exchange, error) {
	return []data.Exchange{}, nil
}
