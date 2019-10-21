package main

import (
	"log"
	"net/http"

	// "github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"./gql"
)

func main() {
	rootQuery := gql.CreateQueryType(nil)
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failes to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
		// Playground: true,
	})
	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
