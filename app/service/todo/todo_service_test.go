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
