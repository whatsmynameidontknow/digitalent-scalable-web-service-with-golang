package middleware

import (
	"context"
	"final-project/helper"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp = helper.NewResponse[any](helper.Authentication)

		token := r.Header.Get("Authorization")
		if token == "" {
			logger.Warn("no token provided")
			resp.Error(helper.ErrNotLoggedIn).Code(http.StatusUnauthorized).Send(w)
			return
		}
		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) != 2 {
			logger.Warn("invalid token format")
			resp.Error(helper.ErrNotLoggedIn).Code(http.StatusUnauthorized).Send(w)
			return
		}
		token = splitToken[1]
		claims, err := helper.VerifyJWT(token)
		if err != nil {
			logger.Error("failed to verify token", "error", err.Error())
			resp.Error(helper.ErrNotLoggedIn).Code(http.StatusUnauthorized).Send(w)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), helper.UserIDKey, claims["sub"]))

		next.ServeHTTP(w, r)
	})
}
