package middleware

import "net/http"

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("panic recovered", "err", err)
				http.Error(w, "there's something wrong on our side", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
