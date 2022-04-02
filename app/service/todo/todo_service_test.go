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

	testutil.RunWithDB(t, name, "default", func(t *testing.T, db *sql.DB) {
		//　予測値
		want := make([]*model.Todo, 3)
		todo1 := model.Todo{
			ID:        strconv.FormatUint(1, 10),
			Text:      "todo1",
			Comment:   "todo1のコメント",
			CreatedAt: TIME_LAYOUT,
			UpdatedAt: TIME_LAYOUT,
		}
		todo2 := model.Todo{
			ID:        strconv.FormatUint(2, 10),
			Text:      "todo2",
			Comment:   "todo2のコメント",
			CreatedAt: TIME_LAYOUT,
			UpdatedAt: TIME_LAYOUT,
		}
		todo3 := model.Todo{
			ID:        strconv.FormatUint(3, 10),
			Text:      "todo3",
			Comment:   "todo3のコメント",
			CreatedAt: TIME_LAYOUT,
			UpdatedAt: TIME_LAYOUT,
		}
		want[0] = &todo1
		want[1] = &todo2
		want[2] = &todo3

		s := &Service{
			db: db,
		}
		//	実行
		result, err := s.TodoList(context.Background())
		if err != nil {
			t.Errorf("%s() error = %v", name, err)
		}

		for i, res := range result {
			if diff := cmp.Diff(*res, *want[i], cmpopts.IgnoreFields(*res, "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
				t.Errorf("%v", diff)
			}
		}

	})
}
