package schema

import (
	"graphql-api/pkg/models"

	"github.com/graphql-go/graphql"
)

// Plant graphql type
var Plant = graphql.NewObject(graphql.ObjectConfig{
	Name: "Plant",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// PlantQueries are the graphql query definitions
var PlantQueries = graphql.Fields{
	"plants": plantsQuery,
}

var plantsQuery = &graphql.Field{
	Type: graphql.NewList(Plant),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return models.Plants, nil
	},
}
