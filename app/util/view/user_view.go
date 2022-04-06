package view

import (
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database/entity"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/model"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/timeutil"
	"strconv"
)

func NewUserFromModel(entity *entity.User) *model.User {
	resUser := model.User{
		ID:        strconv.FormatUint(entity.ID, 10),
		Name:      entity.Name,
		Email:     entity.Email,
		CreatedAt: timeutil.TimeFormat(entity.CreatedAt),
		UpdatedAt: timeutil.TimeFormat(entity.UpdatedAt),
	}

	if entity.ImageURL.Valid {
		imageUrl := entity.ImageURL.String
		resUser.ImageURL = &imageUrl
	}

	if entity.DeletedAt.Valid {
		deletedAt := timeutil.TimeFormat(entity.DeletedAt.Time)
		resUser.DeletedAt = &deletedAt
	}

	return &resUser
}
