package todo

import (
	"context"
	"database/sql"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database/entity"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/model"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"strconv"
)

type Service struct {
	db *sql.DB
}

func LazyInit(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

const TimeLayout = "2006-01-02 15:04:05"

func (s *Service) TodoList(ctx context.Context) ([]*model.Todo, error) {
	todoList, todoErr := entity.Todos().All(ctx, s.db)
	if todoErr != nil {
		return nil, todoErr
	}
	resTodoList := make([]*model.Todo, len(todoList))
	for i, todo := range todoList {
		resTodo := model.Todo{
			ID:        strconv.FormatUint(todo.ID, 10),
			Title:     todo.Title,
			Comment:   todo.Comment,
			CreatedAt: todo.CreatedAt.Format(TimeLayout),
			UpdatedAt: todo.UpdatedAt.Format(TimeLayout),
		}
		if todo.DeletedAt.Valid {
			deletedAt := todo.DeletedAt.Time.Format(TimeLayout)
			resTodo.DeletedAt = &deletedAt
		}
		resTodoList[i] = &resTodo
	}
	return resTodoList, nil
}

func (s *Service) TodoDetail(ctx context.Context, id string) (*model.Todo, error) {
	todo, todoErr := entity.Todos(qm.Where("id=?", id)).One(ctx, s.db)
	if todoErr != nil {
		return nil, todoErr
	}
	resTodo := model.Todo{
		ID:        strconv.FormatUint(todo.ID, 10),
		Title:     todo.Title,
		Comment:   todo.Comment,
		CreatedAt: todo.CreatedAt.Format(TimeLayout),
		UpdatedAt: todo.UpdatedAt.Format(TimeLayout),
	}
	if todo.DeletedAt.Valid {
		deletedAt := todo.DeletedAt.Time.Format(TimeLayout)
		resTodo.DeletedAt = &deletedAt
	}
	return &resTodo, nil
}
