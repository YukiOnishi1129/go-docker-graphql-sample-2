package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/generated"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/service/todo"
	"log"
	"net/http"
)

const containerPort = "3000"

func main() {
	db, dbErr := database.Init()
	if dbErr != nil {
		panic(dbErr)
	}

	todoService := todo.LazyInit(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(todoService)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "4000")
	log.Fatal(http.ListenAndServe(":"+containerPort, nil))

}
