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

	router := chi.NewRouter()

	jwtSecret := os.Getenv("JWT")
	if jwtSecret == "" {
		jwtSecret = "29607b9e17f4c5266a2d33aca075ab62"
	}

	a := auth.New(data.Db, jwtSecret)

	data.Db, _ = data.ConnectPostgres()
	fmt.Println("Connection with database has a successful!")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Use(a.HTTPMiddleware())

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	router.HandleFunc("/login", a.CreateTokenEndpoint)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
