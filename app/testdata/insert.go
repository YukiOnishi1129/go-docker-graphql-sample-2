package testdata

import "database/sql"

var (
	ins *sql.Stmt
	err error
)

func CreateTestData(con *sql.DB) error {
	// todoテーブルのテストデータ作成
	if err = CreateTodoData(con); err != nil {
		return err
	}
	return nil
}
