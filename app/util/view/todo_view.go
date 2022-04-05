package view

import (
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database/entity"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/model"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/timeutil"
	"strconv"
)

func NewTodoFromModel(entity *entity.Todo) *model.Todo {
	resTodo := model.Todo{
		ID:        strconv.FormatUint(entity.ID, 10),
		Title:     entity.Title,
		Comment:   entity.Comment,
		CreatedAt: timeutil.TimeFormat(entity.CreatedAt),
		UpdatedAt: timeutil.TimeFormat(entity.UpdatedAt),
	}
	if entity.DeletedAt.Valid {
		deletedAt := timeutil.TimeFormat(entity.DeletedAt.Time)
		resTodo.DeletedAt = &deletedAt
	}

	return &resTodo
}
