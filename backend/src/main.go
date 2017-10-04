package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/graphql-go/handler"
	"net/http"
	"os"

	"./database"
	"./graphql"
)

func main() {
	fmt.Printf("Serving at port 8080\n")

	// Create a GraphQL handler
	graphqlHandler := handler.New(&handler.Config{
		Schema:   &graphql.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Create Postgres middleware
	options := pg.Options{
		User:     os.Getenv("USER"),
		Database: "crypto-accounts-viewer",
	}
	db := database.Connect(&options)
	handler := database.NewHandler(db, graphqlHandler)

	// Start the Postgres-backed GraphQL server
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}
