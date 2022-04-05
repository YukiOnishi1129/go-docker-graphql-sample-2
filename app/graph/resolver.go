package graph

import (
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/service/todo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todoService *todo.Service
}

func NewResolver(
	todoService *todo.Service,
) *Resolver {
	return &Resolver{
		todoService: todoService,
	}
}
