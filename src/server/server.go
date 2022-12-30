package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-pg/pg/v9"
	src "github.com/pjfcs/goGraphQL_1.13"
	"github.com/pjfcs/goGraphQL_1.13/postgres"
)

const defaultPort = "9000"

func main() {
	DB := postgres.New(&pg.Options{
		User:     "usergestor",
		Password: "1234",
		Database: "meetmeup_dev",
		Addr:     "dbPostgreSQL:5432",
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := src.Config{Resolvers: &src.Resolver{
		MeetupsRepo: postgres.MeetupsRepo{DB: DB},
		UserRepo:    postgres.UsersRepo{DB: DB},
	}}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(src.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
