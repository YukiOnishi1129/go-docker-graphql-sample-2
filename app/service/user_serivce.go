package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database/entity"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/model"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/auth"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/validate"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/util/view"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
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
	// パスワード判定処理
	if input.Password != input.PasswordConfirm {
		return nil, view.NewBadRequestErrorFromModel("パスワードが一致しません。")
	}
	// メールアドレス判定処理
	sameEmailUser, sameEmailUserErr := entity.Users(qm.Where("email=?", input.Email)).One(ctx, s.db)
	if sameEmailUserErr != nil && sameEmailUserErr.Error() != "sql: no rows in result set" {
		return nil, view.NewBadRequestErrorFromModel(sameEmailUserErr.Error())
	}
	if sameEmailUser != nil {
		return nil, view.NewBadRequestErrorFromModel(fmt.Sprintf("メールアドレス「%s」は使用されています。", input.Email))
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
	return "ログアウトしました。", nil
}

// UpdateUserName ユーザー名変更
func (s *UserService) UpdateUserName(ctx context.Context, name string, adminUser *entity.User) (*model.User, error) {
	var err error
	// バリデーション
	if err = validate.UpdateUserNameValidation(validate.UpdateUserNameInput{Name: name}); err != nil {
		return nil, view.NewBadRequestErrorFromModel(err.Error())
	}
	// 更新処理
	adminUser.Name = name
	_, err = adminUser.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, view.NewDBErrorFromModel(err)
	}
	return view.NewUserFromModel(adminUser), nil
}

// UpdateUserEmail ユーザーメールアドレス変更処理
func (s *UserService) UpdateUserEmail(ctx context.Context, email string, adminUser *entity.User) (*model.User, error) {
	var err error
	// バリデーション
	if err = validate.UpdateUserEmailValidation(validate.UpdateUserEmailInput{Email: email}); err != nil {
		return nil, view.NewBadRequestErrorFromModel(err.Error())
	}
	// メールアドレス判定処理
	sameEmailUser, sameEmailUserErr := entity.Users(qm.Where("email=?", email)).One(ctx, s.db)
	if sameEmailUserErr != nil && sameEmailUserErr.Error() != "sql: no rows in result set" {
		return nil, view.NewBadRequestErrorFromModel(sameEmailUserErr.Error())
	}
	if sameEmailUser != nil {
		return nil, view.NewBadRequestErrorFromModel(fmt.Sprintf("メールアドレス「%s」は使用されています。", email))
	}
	// 更新処理
	adminUser.Email = email
	_, err = adminUser.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, view.NewDBErrorFromModel(err)
	}
	return view.NewUserFromModel(adminUser), nil
}

// UpdateUserPassword ユーザーパスワード変更
func (s *UserService) UpdateUserPassword(ctx context.Context, input model.UpdatePasswordInput, adminUser *entity.User) (*model.User, error) {
	var err error
	// バリデーション
	if err = validate.UpdateUserPasswordValidation(input); err != nil {
		return nil, view.NewBadRequestErrorFromModel(err.Error())
	}
	// 現在のパスワード照合処理
	targetUser, targetUserErr := entity.Users(qm.Where("id=?", adminUser.ID)).One(ctx, s.db)
	if targetUserErr != nil {
		return nil, view.NewBadRequestErrorFromModel(targetUserErr.Error())
	}
	if err = bcrypt.CompareHashAndPassword([]byte(targetUser.Password), []byte(input.OldPassword)); err != nil {
		return nil, view.NewUnauthorizedErrorFromModel("現在のパスワードが違います。")
	}
	// パスワード判定処理
	if input.NewPassword != input.NewPasswordConfirm {
		return nil, view.NewBadRequestErrorFromModel("新しいパスワードと確認用パスワードが一致しません。")
	}
	// 更新処理
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	adminUser.Password = string(hashPassword)
	_, err = adminUser.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, view.NewDBErrorFromModel(err)
	}

	return view.NewUserFromModel(adminUser), nil
}

// UploadUserFile ファイルアップロード
func (s *UserService) UploadUserFile(ctx context.Context, file *graphql.Upload, adminUser *entity.User) (*model.User, error) {
	const filePerm fs.FileMode = 0644

	var err error
	// fileデータの読み込み
	targetFile, targetFileErr := io.ReadAll(file.File)
	if targetFileErr != nil {
		return nil, err
	}
	_, fileName, _, _ := runtime.Caller(0)
	filePath := fmt.Sprintf("%s/../assets/images/%s", filepath.Dir(fileName), file.Filename)
	// ファイルがない場合新規作成し、書き込み権限付きで開く
	// https://zenn.dev/hsaki/books/golang-io-package/viewer/file#%E6%9B%B8%E3%81%8D%E8%BE%BC%E3%81%BF%E6%A8%A9%E9%99%90%E4%BB%98%E3%81%8D%E3%81%A7%E9%96%8B%E3%81%8F
	if _, createFileErr := os.Create(filePath); createFileErr != nil {
		return nil, createFileErr
	}
	// ファイルへの書き込み
	// https://zenn.dev/hsaki/books/golang-io-package/viewer/file#%E6%9B%B8%E3%81%8D%E8%BE%BC%E3%81%BF%E6%A8%A9%E9%99%90%E4%BB%98%E3%81%8D%E3%81%A7%E9%96%8B%E3%81%8F
	if err = os.WriteFile(filePath, targetFile, filePerm); err != nil {
		fmt.Println("===========")
		return nil, err
	}
	//pwd, err := os.Getwd()
	//if err != nil {
	//	return nil, err
	//}
	adminUser.ImageURL = null.StringFromPtr(&filePath)
	_, err = adminUser.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, view.NewDBErrorFromModel(err)
	}
	return view.NewUserFromModel(adminUser), nil
}

// MyUserDetail ログインユーザーの詳細情報
func (s *UserService) MyUserDetail(adminUser *entity.User) (*model.User, error) {
	return view.NewUserFromModel(adminUser), nil
}

// UserDetail ユーザー詳細情報
func (s *UserService) UserDetail(ctx context.Context, id string) (*model.User, error) {
	// バリデーション
	if id == "" {
		return nil, view.NewBadRequestErrorFromModel("IDは必須です。")
	}
	user, userErr := entity.Users(qm.Where("id=?", id)).One(ctx, s.db)
	if userErr != nil {
		return nil, view.NewDBErrorFromModel(userErr)
	}
	return view.NewUserFromModel(user), nil
}
