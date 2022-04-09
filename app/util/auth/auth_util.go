package auth

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/database/entity"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
)

type contextKey struct {
	uuid string
}

var userCtxKey = &contextKey{"user"}
var httpWriterKey = &contextKey{"httpWriter"}
var authCookieKey = "auth-cookie"

// MiddleWare cookie認証
func MiddleWare(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			c, authErr := r.Cookie(authCookieKey)
			// put it in context
			ctx := context.WithValue(r.Context(), httpWriterKey, w)

			// and call the next with our new context
			r = r.WithContext(ctx)

			if authErr != nil || c == nil {
				next.ServeHTTP(w, r)
				fmt.Println("=====yyyy=======")
				return
			}

			userId, err := getUserIdFromJwt(c)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			authCtx := context.WithValue(r.Context(), userCtxKey, userId)
			r = r.WithContext(authCtx)

			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) (int, error) {
	userId, err := ctx.Value(userCtxKey).(int)
	if !err {
		return 0, errors.New("認証情報がありません。")
	}
	return userId, nil
}

// RemoveAuthCookie called when user wants to log out, return an instantly expired cookie
func RemoveAuthCookie(ctx context.Context) {
	writer, _ := ctx.Value(httpWriterKey).(http.ResponseWriter)
	http.SetCookie(writer, &http.Cookie{
		HttpOnly: true,
		MaxAge:   0,
		Secure:   true,
		Name:     authCookieKey,
	})
}

// SetAuthCookie can be used inside resolvers to set a cookie
func SetAuthCookie(ctx context.Context, user *entity.User) {
	sessionToken, err := createJwtToken(user)
	if err != nil {
		fmt.Println("Error: create jwt error", err)
		return
	}

	writer, _ := ctx.Value(httpWriterKey).(http.ResponseWriter)

	week := 60 * 60 * 24 * 7
	//week := 3

	cookie := http.Cookie{
		HttpOnly: true,
		MaxAge:   week * 2,
		Secure:   true,
		Name:     authCookieKey,
		Value:    sessionToken,
	}
	http.SetCookie(writer, &cookie)

}

func createJwtToken(user *entity.User) (string, error) {
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = strconv.Itoa(int(user.ID)) + user.Email + user.Name
	claims["id"] = user.ID
	claims["name"] = user.Name
	// latを取り除かないとミドルウェアで「Token used before issued」エラーになる
	// https://github.com/dgrijalva/jwt-go/issues/314#issuecomment-812775567
	// claims["iat"] = time.Now() // jwtの発行時間
	// 経過時間
	// 経過時間を過ぎたjetは処理しないようになる
	// ここでは2習慣の経過時間をリミットにしている
	claims["exp"] = time.Now().Add(time.Hour * 24 * 14).Unix()
	// .envを読み込む
	//err := godotenv.Load()
	//if err != nil {
	//	fmt.Println(err)
	//	return "", err
	//}
	//fmt.Println("%s", os.Getenv("JWT_KEY"))

	// 電子署名
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getUserIdFromJwt(c *http.Cookie) (int, error) {
	clientToken := c.Value
	if clientToken == "" {
		return 0, errors.New("not token")
	}

	secretKey := os.Getenv("JWT_KEY")

	token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("トークンをjwtにparseできません。")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, claimOk := token.Claims.(jwt.MapClaims)
	if !claimOk || !token.Valid {
		return 0, errors.New("id type not match")
	}

	userId, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("id type not match")
	}

	return int(userId), nil
}
