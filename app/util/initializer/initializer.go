package initializer

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/generated"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/service/todo"
)

func Init() (*handler.Server, error) {
	db, dbErr := database.Init()
	if dbErr != nil {
		return nil, dbErr
	}

	todoService := todo.LazyInit(db)

	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(todoService)})), nil
}
