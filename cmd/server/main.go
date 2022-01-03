package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "hello rodler!", nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func main() {
	// Create a graphql-go HTTP hanler with our previously defined schema
	// and we also set it to return pretty JSON output

	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)

	http.ListenAndServe(":8080", nil)
}
