package schema

import (
	"log"

	"github.com/graphql-go/graphql"
)

// NewSchema creates a graphql schema
func NewSchema() graphql.Schema {
	queryFields := []graphql.Fields{
		PlantQueries,
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: mergeFields(queryFields),
		}),
	})
	if err != nil {
		log.Fatal(err)
	}

	return schema
}

func mergeFields(fieldsList []graphql.Fields) graphql.Fields {
	result := make(map[string]*graphql.Field)

	for _, f := range fieldsList {
		for k, v := range f {
			result[k] = v
		}
	}

	return result
}
