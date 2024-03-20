package middleware

import (
	"final-project/helper"
	"net/http"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp helper.Response[any]
		defer func() {
			if r := recover(); r != nil {
				resp.Error(helper.ErrInternal).Code(http.StatusInternalServerError).Send(w)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
