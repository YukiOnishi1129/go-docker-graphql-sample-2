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

func TestService_TodoList_OnSuccess(t *testing.T) {

	name := "TodoList"

	testutil.RunWithDB(t, "TodoList get", func(t *testing.T, db *sql.DB) {
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
			t.Errorf("%s() error = %v", name, err)
		}

		// テスト結果の評価
		for i, res := range result {
			if diff := cmp.Diff(*res, *want[i], cmpopts.IgnoreFields(*res, "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
				t.Errorf("%v", diff)
			}
		}

	})
}
