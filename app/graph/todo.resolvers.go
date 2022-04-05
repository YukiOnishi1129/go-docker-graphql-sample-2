package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/generated"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	return r.todoService.CreateTodo(ctx, input)
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	return r.todoService.UpdateTodo(ctx, input)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (string, error) {
	return r.todoService.DeleteTodo(ctx, id)
}

func (r *queryResolver) TodoList(ctx context.Context) ([]*model.Todo, error) {
	return r.todoService.TodoList(ctx)
}

func (r *queryResolver) TodoDetail(ctx context.Context, id string) (*model.Todo, error) {
	return r.todoService.TodoDetail(ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
