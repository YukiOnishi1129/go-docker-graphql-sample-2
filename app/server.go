package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/generated"
	"log"
	"net/http"
)

const containerPort = "3000"

func main() {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "4000")
	log.Fatal(http.ListenAndServe(":"+containerPort, nil))

	//fmt.Println("server start")
	//router := mux.NewRouter().StrictSlash(true)
	//
	//db, dbErr := database.Init()
	//if dbErr != nil {
	//	panic(dbErr)
	//}
	//
	//ctx := context.Background()
	//todo, todoErr := models.Todos(
	//	qm.Where("id=?", 2),
	//).One(ctx, db)
	//if todoErr != nil {
	//	panic(todoErr)
	//}
	//
	//fmt.Println(todo.ID)
	//
	//if err := http.ListenAndServe(fmt.Sprintf(":%d", 3000), router); err != nil {
	//	panic(err)
	//}
}
