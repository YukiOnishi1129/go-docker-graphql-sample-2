package auth

import (
	"context"
	"fmt"
	"net/http"
)

type contextKey struct {
	uuid string
}

var userCtxKey = &contextKey{"user"}
var httpWriterKey = &contextKey{"httpWriter"}

//CookieMiddleWare
func CookieMiddleWare() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// c, err := r.Cookie("session-cookie")

			// put it in context
			ctx := context.WithValue(r.Context(), httpWriterKey, w)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

//RemoveAuthCookie called when user wants to log out, return an instantly expired cookie
func RemoveAuthCookie(ctx context.Context) {
	writer, _ := ctx.Value(httpWriterKey).(http.ResponseWriter)
	http.SetCookie(writer, &http.Cookie{
		HttpOnly: true,
		MaxAge:   0,
		Secure:   true,
		Name:     "auth-cookie",
	})
}

//SetAuthCookie can be used inside resolvers to set a cookie
func SetAuthCookie(ctx context.Context, sessionToken string) {

	writer, _ := ctx.Value(httpWriterKey).(http.ResponseWriter)

	week := 60 * 60 * 24 * 7

	cookie := http.Cookie{
		HttpOnly: true,
		MaxAge:   week * 2,
		Secure:   true,
		Name:     "auth-cookie",
		Value:    sessionToken,
	}
	fmt.Println("===========%v", writer)
	http.SetCookie(writer, &cookie)

}
