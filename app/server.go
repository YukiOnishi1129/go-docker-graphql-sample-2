package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/initializer"
	"log"
	"net/http"
)

const containerPort = "3000"

func main() {
	srv, err := initializer.Init()
	if err != nil {
		panic(err)
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "4000")
	log.Fatal(http.ListenAndServe(":"+containerPort, nil))

}
