package graphql

import (
	"github.com/graphql-go/graphql"

	"../database"
)

// Exported schema, created in init()
var Schema graphql.Schema

func init() {
	// Currencies
	currency := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Currency",
		Description: "A currency",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.Int},
			"name":     &graphql.Field{Type: graphql.String},
			"longName": &graphql.Field{Type: graphql.String},
		},
	})
	currencyList := graphql.NewList(currency)

	// Exchanges
	exchange := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Exchange",
		Description: "A crypto-currency exchange",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
		},
	})
	exchangeList := graphql.NewList(exchange)

	// Accounts
	account := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Account",
		Description: "A crypto-currency account on an exchange",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.Int},
			"currency": &graphql.Field{Type: currency},
			"exchange": &graphql.Field{Type: exchange},
		},
	})
	accountList := graphql.NewList(account)

	// Queries
	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"accounts": &graphql.Field{
				Type: accountList,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db := database.GetDatabaseFromContext(params.Context)
					return database.GetAccounts(db)
				},
			},
			"currencies": &graphql.Field{
				Type: currencyList,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db := database.GetDatabaseFromContext(params.Context)
					return database.GetCurrencies(db)
				},
			},
			"exchanges": &graphql.Field{
				Type: exchangeList,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db := database.GetDatabaseFromContext(params.Context)
					return database.GetExchanges(db)
				},
			},
		},
	})

	// Account creation
	createAccount := graphql.Field{
		Type: account,
		Args: graphql.FieldConfigArgument{
			"name":       &graphql.ArgumentConfig{Type: graphql.String},
			"currencyId": &graphql.ArgumentConfig{Type: graphql.Int},
			"exchangeId": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			db := database.GetDatabaseFromContext(params.Context)
			return database.CreateAccount(db, params.Args)
		},
	}

	// Mutations
	mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createAccount": &createAccount,
		},
	})

	// Create the schema from our query and mutation types
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})

	// Panic if there is a problem with the schema
	if err != nil {
		panic(err)
	}
}
