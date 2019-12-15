package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/mmcdole/gofeed"
)

func SchemaHandler() http.Handler {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"parse": &graphql.Field{
				Args: graphql.FieldConfigArgument{
					"rssUrl": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Type: Feed,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rssURL := p.Args["rssUrl"].(string)

					fp := gofeed.NewParser()
					feed, _ := fp.ParseURL(rssURL)

					return feed, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})

	return handler.New(&handler.Config{Schema: &schema, Pretty: true, GraphiQL: true})
}
