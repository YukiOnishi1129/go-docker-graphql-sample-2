package testdata

import (
	"database/sql"
)

type createTodoData struct {
	ID      int64
	Title   string
	Comment string
}

func CreateTodoData(con *sql.DB) error {
	insertDataList := [...]*createTodoData{
		{
			ID:      1,
			Title:   "todo1",
			Comment: "todo1のコメント",
		},
		{
			ID:      2,
			Title:   "todo2",
			Comment: "todo2のコメント",
		},
		{
			ID:      3,
			Title:   "todo3",
			Comment: "todo3のコメント",
		},
	}

	for _, insertData := range insertDataList {
		ins, err = con.Prepare("INSERT INTO todos (id, title, comment) VALUES (?,?,?)")
		if err != nil {
			return err
		}
		_, err = ins.Exec(insertData.ID, insertData.Title, insertData.Comment)
		if err != nil {
			return err
		}
	}
	return nil
}
