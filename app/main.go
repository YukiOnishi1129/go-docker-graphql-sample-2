package main

import (
	"context"
	"fmt"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database"
	models "github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/model"
	"github.com/gorilla/mux"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
)

func main() {
	fmt.Println("server start")
	router := mux.NewRouter().StrictSlash(true)

	db, dbErr := database.Init()
	if dbErr != nil {
		return
	}

	ctx := context.Background()
	todo, todoErr := models.Todos(
		qm.Where("id=?", 2),
	).One(ctx, db)
	if todoErr != nil {
		return
	}

	fmt.Println(todo.ID)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", 3000), router); err != nil {
		return
	}
}
