package todo

import (
	"context"
	"database/sql"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/model"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/testutil"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"strconv"
	"testing"
)

const TimeLayout = "2006-01-02 15:04:05"

func TestService_TodoList_OnSuccess(t *testing.T) {
	testutil.RunWithDB(t, "get TodoList", func(t *testing.T, db *sql.DB) {
		//　予測値
		want := [...]*model.Todo{
			{
				ID:        strconv.FormatUint(1, 10),
				Title:     "todo1",
				Comment:   "todo1のコメント",
				CreatedAt: TimeLayout,
				UpdatedAt: TimeLayout,
			},
			{
				ID:        strconv.FormatUint(2, 10),
				Title:     "todo2",
				Comment:   "todo2のコメント",
				CreatedAt: TimeLayout,
				UpdatedAt: TimeLayout,
			},
			{
				ID:        strconv.FormatUint(3, 10),
				Title:     "todo3",
				Comment:   "todo3のコメント",
				CreatedAt: TimeLayout,
				UpdatedAt: TimeLayout,
			},
		}

		s := &Service{
			db: db,
		}
		//	実行
		result, err := s.TodoList(context.Background())
		if err != nil {
			t.Errorf("get TodoList() error = %v", err)
		}

		// テスト結果の評価
		for i, res := range result {
			if diff := cmp.Diff(*res, *want[i], cmpopts.IgnoreFields(*res, "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
				t.Errorf("%v", diff)
			}
		}
	})
}

func TestService_TodoDetail_OnSuccess(t *testing.T) {
	testutil.RunWithDB(t, "get TodoDetail", func(t *testing.T, db *sql.DB) {
		//　予測値
		want := model.Todo{
			ID:        strconv.FormatUint(2, 10),
			Title:     "todo2",
			Comment:   "todo2のコメント",
			CreatedAt: TimeLayout,
			UpdatedAt: TimeLayout,
		}

		s := &Service{
			db: db,
		}
		targetId := 2
		//	実行
		result, err := s.TodoDetail(context.Background(), strconv.Itoa(targetId))
		if err != nil {
			t.Errorf("get TodoDetail() error = %v", err)
		}

		if diff := cmp.Diff(*result, want, cmpopts.IgnoreFields(*result, "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
			t.Errorf("%v", diff)
		}
	})
}

func TestService_TodoDetail_OnFailure(t *testing.T) {
	testutil.RunWithDB(t, "get TodoDetail error", func(t *testing.T, db *sql.DB) {
		s := &Service{
			db: db,
		}
		targetId := 4
		//	実行
		result, err := s.TodoDetail(context.Background(), strconv.Itoa(targetId))

		if err == nil {
			t.Fatalf("存在しないtodoはエラーになるべきです. err: %v", err)
		}
		if result != nil {
			t.Errorf("nilであるべきです. got: %v", result)
		}
	})
}

func TestService_CreateTodo_OnSuccess(t *testing.T) {
	testutil.RunWithDB(t, "create todo success", func(t *testing.T, db *sql.DB) {
		// 予測値
		want := model.Todo{
			ID:        strconv.FormatUint(4, 10),
			Title:     "todo4",
			Comment:   "todo4のコメント",
			CreatedAt: TimeLayout,
			UpdatedAt: TimeLayout,
		}
		s := &Service{
			db: db,
		}
		args := model.CreateTodoInput{
			Title:   "todo4",
			Comment: "todo4のコメント",
		}
		//	実行
		result, err := s.CreateTodo(context.Background(), args)
		if err != nil {
			t.Errorf("CreateTodo() error = %v", err)
		}
		if diff := cmp.Diff(*result, want, cmpopts.IgnoreFields(*result, "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
			t.Errorf("%v", diff)
		}
	})
}

func TestService_CreateTodo_OnFailure(t *testing.T) {
	testutil.RunWithDB(t, "create todo bad request empty title", func(t *testing.T, db *sql.DB) {
		// 予測値
		s := &Service{
			db: db,
		}
		args := model.CreateTodoInput{
			Title:   "",
			Comment: "todo4のコメント",
		}
		//	実行
		result, err := s.CreateTodo(context.Background(), args)

		if err == nil {
			t.Fatalf("titleのバリデーションエラーになるべきです. err: %v", err)
		}
		if result != nil {
			t.Errorf("nilであるべきです. got: %v", result)
		}
	})

	testutil.RunWithDB(t, "create todo bad request empty comment", func(t *testing.T, db *sql.DB) {
		// 予測値
		s := &Service{
			db: db,
		}
		args := model.CreateTodoInput{
			Title:   "todo4のタイトル",
			Comment: "",
		}
		//	実行
		result, err := s.CreateTodo(context.Background(), args)

		if err == nil {
			t.Fatalf("commentのバリデーションエラーになるべきです. err: %v", err)
		}
		if result != nil {
			t.Errorf("nilであるべきです. got: %v", result)
		}
	})
}

func TestService_UpdateTodo_OnSuccess(t *testing.T) {
	testutil.RunWithDB(t, "update todo ", func(t *testing.T, db *sql.DB) {
		// 予測値
		want := model.Todo{
			ID:        strconv.FormatUint(3, 10),
			Title:     "todo3title",
			Comment:   "todo3コメントupdate",
			CreatedAt: TimeLayout,
			UpdatedAt: TimeLayout,
		}
		s := &Service{
			db: db,
		}
		args := model.UpdateTodoInput{
			ID:      "3",
			Title:   "todo3title",
			Comment: "todo3コメントupdate",
		}
		//	実行
		result, err := s.UpdateTodo(context.Background(), args)
		if err != nil {
			t.Errorf("UpdateTodo() error = %v", err)
		}
		if diff := cmp.Diff(*result, want, cmpopts.IgnoreFields(*result, "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
			t.Errorf("%v", diff)
		}
	})
}

func TestService_UpdateTodo_OnFailure(t *testing.T) {
	testutil.RunWithDB(t, "update todo bad request title empty", func(t *testing.T, db *sql.DB) {
		s := &Service{
			db: db,
		}
		args := model.UpdateTodoInput{
			ID:      "3",
			Title:   "",
			Comment: "todo3コメントupdate",
		}
		//	実行
		result, err := s.UpdateTodo(context.Background(), args)

		if err == nil {
			t.Fatalf("idのバリデーションエラーになるべきです. err: %v", err)
		}
		if result != nil {
			t.Errorf("nilであるべきです. got: %v", result)
		}
	})

	testutil.RunWithDB(t, "update todo not found", func(t *testing.T, db *sql.DB) {
		s := &Service{
			db: db,
		}
		args := model.UpdateTodoInput{
			ID:      "4",
			Title:   "todo3title",
			Comment: "todo3タイトルupdate",
		}
		//	実行
		result, err := s.UpdateTodo(context.Background(), args)

		if err == nil {
			t.Fatalf("該当データなしでエラーになるべきです. err: %v", err)
		}
		if result != nil {
			t.Errorf("nilであるべきです. got: %v", result)
		}
	})
}

func TestService_DeleteTodo_OnSuccess(t *testing.T) {
	testutil.RunWithDB(t, "delete todo ", func(t *testing.T, db *sql.DB) {
		// 予測値
		want := "1"
		s := &Service{
			db: db,
		}
		args := "1"
		//	実行
		result, err := s.DeleteTodo(context.Background(), args)
		if err != nil {
			t.Errorf("UpdateTodo() error = %v", err)
		}
		if diff := cmp.Diff(result, want); diff != "" {
			t.Errorf("%v", diff)
		}
	})
}

func TestService_DeleteTodo_OnFailure(t *testing.T) {
	testutil.RunWithDB(t, "delete todo not empty id", func(t *testing.T, db *sql.DB) {
		s := &Service{
			db: db,
		}
		args := ""
		//	実行
		result, err := s.DeleteTodo(context.Background(), args)
		if err == nil {
			t.Fatalf("idのバリデーションエラーになるべきです. err: %v", err)
		}
		if result != "" {
			t.Errorf("空文字であるべきです. got: %v", result)
		}
	})

	testutil.RunWithDB(t, "delete todo bad not found", func(t *testing.T, db *sql.DB) {
		s := &Service{
			db: db,
		}
		args := "4"
		//	実行
		result, err := s.DeleteTodo(context.Background(), args)
		if err == nil {
			t.Fatalf("該当データなしでエラーになるべきです. err: %v", err)
		}
		if result != "" {
			t.Errorf("空文字であるべきです. got: %v", result)
		}
	})
}
