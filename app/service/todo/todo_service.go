package todo

import (
	"context"
	"database/sql"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database/entity"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/model"
	validate "github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/validate"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/view"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Service struct {
	db *sql.DB
}

func LazyInit(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

var err error

func (s *Service) TodoList(ctx context.Context) ([]*model.Todo, error) {
	todoList, todoErr := entity.Todos().All(ctx, s.db)
	if todoErr != nil {
		return nil, todoErr
	}
	resTodoList := make([]*model.Todo, len(todoList))
	for i, todo := range todoList {
		resTodoList[i] = view.NewTodoFromModel(todo)
	}
	return resTodoList, nil
}

func (s *Service) TodoDetail(ctx context.Context, id string) (*model.Todo, error) {
	todo, todoErr := entity.Todos(qm.Where("id=?", id)).One(ctx, s.db)
	if todoErr != nil {
		return nil, todoErr
	}
	return view.NewTodoFromModel(todo), nil
}

func (s *Service) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	// バリデーション
	if err = validate.CreateTodoValidation(input); err != nil {
		return nil, err
	}

	// 新規登録処理
	newTodo := &entity.Todo{
		Title:   input.Title,
		Comment: input.Comment,
	}
	if err = newTodo.Insert(ctx, s.db, boil.Infer()); err != nil {
		return nil, err
	}

	return view.NewTodoFromModel(newTodo), nil
}

func (s *Service) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	// バリデーション
	if err = validate.UpdateTodoValidation(input); err != nil {
		return nil, err
	}
	todo, todoErr := entity.Todos(qm.Where("id=?", input.ID)).One(ctx, s.db)
	if todoErr != nil {
		return nil, todoErr
	}

	todo.Title = input.Title
	todo.Comment = input.Comment

	_, updateTodoErr := todo.Update(ctx, s.db, boil.Infer())
	if updateTodoErr != nil {
		return nil, err
	}
	return view.NewTodoFromModel(todo), nil
}
