package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {

	rootQuery := graphql.ObjectConfig{
		Name: "RootQuery",

		// Schema
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fmt.Printf("could be anything inside here!\n")
					return "world", nil
				},
			},
		},
	}

	schema, e := graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)})
	if e != nil {
		log.Fatalf("Failed to create new schema, error: %v", e)
	}

	query := `
		{
			hello
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}

	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("Failed to execute graphql operation, errors: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)

	fmt.Printf("%s \n", rJSON)
}
