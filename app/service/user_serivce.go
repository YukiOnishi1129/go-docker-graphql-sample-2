package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database/entity"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/model"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/auth"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/validate"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/view"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

// SignIn ログイン
func (s *UserService) SignIn(ctx context.Context, input model.SignInInput) (*model.User, error) {
	var err error
	// バリデーション
	if err = validate.SignInValidation(input); err != nil {
		return nil, view.NewBadRequestErrorFromModel(err.Error())
	}

	// ユーザー認証
	user, err := entity.Users(qm.Where("email=?", input.Email)).One(ctx, s.db)
	if err != nil {
		return nil, view.NewDBErrorFromModel(err)
	}
	// パスワード照合
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, view.NewUnauthorizedErrorFromModel("パスワードが違います。")
	}

	// cookieに保持
	auth.SetAuthCookie(ctx, user)
	return view.NewUserFromModel(user), nil
}

// SignUp 会員登録
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

	// cookieに保持
	auth.SetAuthCookie(ctx, newUser)

	return view.NewUserFromModel(newUser), nil
}

// SignOut ログアウト
func (s *UserService) SignOut(ctx context.Context) (string, error) {
	auth.RemoveAuthCookie(ctx)
	return fmt.Sprintf("ログアウトしました。"), nil
}
