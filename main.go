package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/graphql-go/graphql"

	_ "github.com/lib/pq"
)

type Plant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var plants = []Plant{
	Plant{
		ID:   "001",
		Name: "Fiddle Leaf Fig",
	},
	Plant{
		ID:   "002",
		Name: "Swiss Cheese Plant",
	},
	Plant{
		ID:   "003",
		Name: "Macho Fern",
	},
	Plant{
		ID:   "004",
		Name: "ZZ Plant",
	},
}

func main() {
	// connectToDB()

	http.Handle("/", handleGraphQLPlayground())
	http.Handle("/api", handleGraphQL())

	fmt.Println("listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

func connectToDB() *sql.DB {
	user := os.Getenv("POSTGRES_USER")
	pw := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=db sslmode=disable",
		user,
		pw,
		dbName,
	)
	fmt.Println(user)
	fmt.Println(pw)
	fmt.Println(dbName)
	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func handleGraphQLPlayground() http.HandlerFunc {
	type playgroundData struct {
		Endpoint string
	}

	tpl := template.Must(template.ParseFiles("./graphql-playground.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, playgroundData{
			Endpoint: "http://localhost:3000/api",
		})
	}
}

func handleGraphQL() http.HandlerFunc {
	plantType := graphql.NewObject(graphql.ObjectConfig{
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

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"plants": &graphql.Field{
					Type: graphql.NewList(plantType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return plants, nil
					},
				},
			},
		}),
	})

	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  params.Query,
			OperationName:  params.OperationName,
			VariableValues: params.Variables,
		})
		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
	}
}
