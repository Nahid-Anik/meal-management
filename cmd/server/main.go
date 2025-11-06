package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/nahid/meal-management/graph"
	"github.com/nahid/meal-management/graph/generated"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	r.Handle("/query", srv)

	log.Printf("🚀 Server running at http://localhost:%s/query", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
