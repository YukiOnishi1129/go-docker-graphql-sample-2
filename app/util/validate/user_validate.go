package validate

import (
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/graph/model"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UpdateUserNameInput struct {
	Name string `json:"name"`
}

// SignUpValidation ログインパラメータのバリデーション
func SignUpValidation(input model.SignUpInput) error {
	return validation.ValidateStruct(&input,
		validation.Field(
			&input.Name,
			validation.Required.Error("お名前は必須入力です。"),
			validation.RuneLength(1, 15).Error("お名前は 1～15 文字です"),
		),
		validation.Field(
			&input.Email,
			validation.Required.Error("メールアドレスは必須入力です"),
			validation.RuneLength(5, 40).Error("メールアドレスは 5～40 文字です"),
			is.Email.Error("メールアドレスを入力して下さい"),
		),
		validation.Field(
			&input.Password,
			validation.Required.Error("パスワードは必須入力です"),
			validation.Length(6, 20).Error("パスワードは6文字以上、20字以内で入力してください。"),
			is.Alphanumeric.Error("パスワードは英数字で入力してください。"),
		),
		validation.Field(
			&input.PasswordConfirm,
			validation.Required.Error("パスワード(確認用)は必須入力です"),
		),
	)
}

// SignInValidation 会員登録パラメータのバリデーション
func SignInValidation(input model.SignInInput) error {
	return validation.ValidateStruct(&input,
		validation.Field(
			&input.Email,
			validation.Required.Error("メールアドレスは必須入力です"),
			validation.RuneLength(5, 40).Error("メールアドレスは 5～40 文字です"),
			is.Email.Error("メールアドレスを入力して下さい"),
		),
		validation.Field(
			&input.Password,
			validation.Required.Error("パスワードは必須入力です"),
			validation.Length(6, 20).Error("パスワードは6文字以上、20字以内で入力してください。"),
			is.Alphanumeric.Error("パスワードは英数字で入力してください。"),
		),
	)
}

// UpdateUserNameValidation ユーザー名変更パラメータのバリデーション
func UpdateUserNameValidation(input UpdateUserNameInput) error {
	return validation.ValidateStruct(&input,
		validation.Field(
			&input.Name,
			validation.Required.Error("お名前は必須入力です。"),
			validation.RuneLength(1, 15).Error("お名前は 1～15 文字です"),
		),
	)
}
