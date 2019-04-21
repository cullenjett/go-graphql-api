package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	_ "github.com/lib/pq"
)

type query struct{}

func (_ *query) Hello() string {
	return "Hello, world!"
}

func main() {
	// db := connectToDB()

	tpl := template.Must(template.ParseFiles("./graphql-playground.html"))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}))

	s := `
		schema {
			query: Query
		}
		type Query {
			hello: String!
		}
	`
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/api", &relay.Handler{Schema: schema})

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
