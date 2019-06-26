package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"graphql-api/pkg/schema"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/graphql-go/graphql"

	_ "github.com/lib/pq"
)

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
	schema := schema.NewSchema()

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
