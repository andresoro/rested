package main

import (
	"log"
	"net/http"

	gql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (*query) Hello() string { return "Hello World" }

func main() {
	s := `
		schema {
			query: Query
		}
		type Query {
			hello: String!
		}
	`

	schema := gql.MustParseSchema(s, &query{})
	http.Handle("/", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}