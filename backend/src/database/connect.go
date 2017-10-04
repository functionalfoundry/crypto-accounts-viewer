package database

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"../data"
)

func newSchema(db *pg.DB) error {
	for _, model := range []interface{}{
		&data.Account{},
		&data.Currency{},
		&data.Exchange{},
		&data.Transfer{},
	} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func Connect(options *pg.Options) *pg.DB {
	// Connect to the database
	db := pg.Connect(options)

	// Create the database tables
	err := newSchema(db)
	if err != nil {
		panic(err)
	}

	// Populate the database with seed data
	err = db.Insert(&data.Currencies)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err = db.Insert(&data.Exchanges)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return db
}
