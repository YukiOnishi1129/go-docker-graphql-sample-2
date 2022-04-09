package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database/entity"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/model"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/auth"
	validate "github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/validate"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/view"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TodoService struct {
	db *sql.DB
}

func LazyInitTodoService(db *sql.DB) *TodoService {
	return &TodoService{
		db: db,
	}
}

func (s *TodoService) TodoList(ctx context.Context) ([]*model.Todo, error) {
	userId, err := auth.ForContext(ctx)
	if err != nil {
		return nil, view.NewUnauthorizedErrorFromModel("認証情報がありません。")
	}
	fmt.Printf("userID: %d", userId)
	todoList, todoErr := entity.Todos().All(ctx, s.db)
	if todoErr != nil {
		return nil, view.NewDBErrorFromModel(todoErr)
	}
	resTodoList := make([]*model.Todo, len(todoList))
	for i, todo := range todoList {
		resTodoList[i] = view.NewTodoFromModel(todo)
	}
	return resTodoList, nil
}

func (s *TodoService) TodoDetail(ctx context.Context, id string) (*model.Todo, error) {
	// バリデーション
	if id == "" {
		return nil, view.NewBadRequestErrorFromModel("IDは必須です。")
	}
	todo, todoErr := entity.Todos(qm.Where("id=?", id)).One(ctx, s.db)
	if todoErr != nil {
		return nil, view.NewDBErrorFromModel(todoErr)
	}
	return view.NewTodoFromModel(todo), nil
}

func (s *TodoService) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	var err error
	// バリデーション
	if err = validate.CreateTodoValidation(input); err != nil {
		return nil, view.NewBadRequestErrorFromModel(err.Error())
	}

	// 新規登録処理
	newTodo := &entity.Todo{
		Title:   input.Title,
		Comment: input.Comment,
	}
	if err = newTodo.Insert(ctx, s.db, boil.Infer()); err != nil {
		return nil, view.NewDBErrorFromModel(err)
	}

	return view.NewTodoFromModel(newTodo), nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	var err error
	// バリデーション
	if err = validate.UpdateTodoValidation(input); err != nil {
		return nil, view.NewBadRequestErrorFromModel(err.Error())
	}
	todo, todoErr := entity.Todos(qm.Where("id=?", input.ID)).One(ctx, s.db)
	if todoErr != nil {
		return nil, view.NewDBErrorFromModel(todoErr)
	}

	// 更新処理
	todo.Title = input.Title
	todo.Comment = input.Comment
	_, err = todo.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, view.NewDBErrorFromModel(err)
	}
	return view.NewTodoFromModel(todo), nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, id string) (string, error) {
	var err error
	// バリデーション
	if id == "" {
		return "", view.NewBadRequestErrorFromModel("IDは必須です。")
	}
	todo, todoErr := entity.Todos(qm.Where("id=?", id)).One(ctx, s.db)
	if todoErr != nil {
		return "", view.NewDBErrorFromModel(todoErr)
	}

	// 削除処置
	if _, err = todo.Delete(ctx, s.db); err != nil {
		return "", view.NewInternalServerErrorFromModel(todoErr.Error())
	}
	return id, nil
}
