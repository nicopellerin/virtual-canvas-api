package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/nicopellerin/virtual-canvas-api/graph"
	"github.com/nicopellerin/virtual-canvas-api/graph/auth"
	"github.com/nicopellerin/virtual-canvas-api/graph/database"
	"github.com/nicopellerin/virtual-canvas-api/graph/generated"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	db := database.New()

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowedHeaders:   []string{"Accept", "Token", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	router.Use(auth.Middleware(db))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{UsersCollection: database.UsersCollection{DB: db}}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
