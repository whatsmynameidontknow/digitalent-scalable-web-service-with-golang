package middleware

import (
	"final-project/helper"
	"net/http"
)

func NewContentTypeMiddleware(allowedContentTypes map[string]struct{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			contentType := r.Header.Get("Content-Type")
			if _, ok := allowedContentTypes[contentType]; !ok {
				var resp = helper.NewResponse[any](helper.Default)
				resp.Error(helper.ErrInvalidContentType).Code(http.StatusUnsupportedMediaType).Send(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

}
