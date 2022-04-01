package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database/entity"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/generated"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/model"
	"strconv"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TodoList(ctx context.Context) ([]*model.Todo, error) {
	todoList, todoErr := entity.Todos().All(ctx, r.DB)
	if todoErr != nil {
		return nil, todoErr
	}
	resTodoList := make([]*model.Todo, len(todoList))
	for i, todo := range todoList {
		resTodo := model.Todo{
			ID:        strconv.FormatUint(todo.ID, 10),
			Text:      todo.Title,
			Comment:   todo.Comment,
			CreatedAt: todo.CreatedAt.Format(TIME_LAYOUT),
			UpdatedAt: todo.UpdatedAt.Format(TIME_LAYOUT),
		}
		if todo.DeletedAt.Valid {
			deletedAt := todo.DeletedAt.Time.Format(TIME_LAYOUT)
			resTodo.DeletedAt = &deletedAt
		}
		resTodoList[i] = &resTodo
	}
	return resTodoList, nil
}

func (r *queryResolver) TodoDetail(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
