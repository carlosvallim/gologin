package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/carlosvallim/gologin/auth"
	"github.com/carlosvallim/gologin/data"
	"github.com/carlosvallim/gologin/graph"
	"github.com/carlosvallim/gologin/graph/generated"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	a := auth.New(data.Db)

	data.Db, _ = data.ConnectPostgres()
	fmt.Println("Connection with database has a successful!")

	router := chi.NewRouter()
	router.Use(a.HTTPMiddleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.HandleFunc("/login", a.CreateTokenEndpoint)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
