package graphql

import (
	"github.com/graphql-go/graphql"

	"../database"
)

// Exported schema, created in init()
var Schema graphql.Schema

func init() {
	// Currencies
	currencyType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Currency",
		Description: "A currency",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
		},
	})
	currencyListType := graphql.NewList(currencyType)

	// Exchanges
	exchangeType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Exchange",
		Description: "A crypto-currency exchange",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
		},
	})
	exchangeListType := graphql.NewList(exchangeType)

	// Accounts
	accountType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Account",
		Description: "A crypto-currency account on an exchange",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.Int},
			"currency": &graphql.Field{Type: currencyType},
			"exchange": &graphql.Field{Type: exchangeType},
		},
	})
	accountListType := graphql.NewList(accountType)

	// Queries
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"accounts": &graphql.Field{
				Type: accountListType,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db := database.GetDatabaseFromContext(params.Context)
					return database.GetAccounts(db)
				},
			},
			"currencies": &graphql.Field{
				Type: currencyListType,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db := database.GetDatabaseFromContext(params.Context)
					return database.GetCurrencies(db)
				},
			},
			"exchanges": &graphql.Field{
				Type: exchangeListType,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db := database.GetDatabaseFromContext(params.Context)
					return database.GetExchanges(db)
				},
			},
		},
	})

	// Mutations
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createAccount": &graphql.Field{
				Type: accountType,
				Args: graphql.FieldConfigArgument{
					"currencyId": &graphql.ArgumentConfig{Type: graphql.Int},
					"exchangeId": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db := database.GetDatabaseFromContext(params.Context)
					return database.CreateAccount(db, params.Args)
				},
			},
		},
	})

	// Create the schema from our query and mutation types
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	// Panic if there is a problem with the schema
	if err != nil {
		panic(err)
	}
}
