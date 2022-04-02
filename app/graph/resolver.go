package graph

import (
	"database/sql"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/service/todo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db              *sql.DB
	todoServiceFunc todo.LazyInitFunc
}

func NewResolver(
	db *sql.DB,
	todoServiceFunc todo.LazyInitFunc,
) *Resolver {
	return &Resolver{
		db:              db,
		todoServiceFunc: todoServiceFunc,
	}
}
