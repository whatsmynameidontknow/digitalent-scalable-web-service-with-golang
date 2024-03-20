package middleware

import (
	"context"
	"final-project/helper"
	"net/http"
	"strings"
)

type ContextKey string

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp helper.Response[any]

		token := r.Header.Get("Authorization")
		if token == "" {
			resp.Error(helper.ErrNotLoggedIn).Code(http.StatusUnauthorized).Send(w)
			return
		}
		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) != 2 {
			resp.Error(helper.ErrNotLoggedIn).Code(http.StatusUnauthorized).Send(w)
			return
		}
		token = splitToken[1]
		claims, err := helper.VerifyJWT(token)
		if err != nil {
			resp.Error(helper.ErrNotLoggedIn).Code(http.StatusUnauthorized).Send(w)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), helper.UserIDKey, claims["sub"]))

		next.ServeHTTP(w, r)
	})
}
