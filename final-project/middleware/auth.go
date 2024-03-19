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
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token = splitToken[1]
		claims, err := helper.VerifyJWT(token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), helper.UserIDKey, claims["sub"]))

		next.ServeHTTP(w, r)
	})
}
