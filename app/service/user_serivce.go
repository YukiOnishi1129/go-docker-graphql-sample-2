package service

import (
	"context"
	"database/sql"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database/entity"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/model"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/validate"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/view"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *sql.DB
}

func LazyInitUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) SignUp(ctx context.Context, input model.SignUpInput) (*model.User, error) {
	var err error
	// バリデーション
	if err = validate.SignUpValidation(input); err != nil {
		return nil, view.NewBadRequestErrorFromModel(err.Error())
	}

	// パスワードハッシュ化
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// DB登録
	newUser := &entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashPassword),
	}

	if err = newUser.Insert(ctx, s.db, boil.Infer()); err != nil {
		return nil, view.NewDBErrorFromModel(err)
	}

	// sessionに保持

	// cookieに保持

	return view.NewUserFromModel(newUser), nil
}
