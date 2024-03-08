package helper

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Response[T any] struct {
	S    bool     `json:"success"`
	E    []string `json:"errors"`
	D    T        `json:"data"`
	code int
}

type Sender interface {
	Send(http.ResponseWriter) error
}

func (r *Response[T]) Data(data T) *Response[T] {
	r.D = data
	return r
}

func (r *Response[T]) Success(ok bool) *Response[T] {
	r.S = ok
	return r
}

func (r *Response[T]) Error(err error) *Response[T] {
	r.E = strings.Split(err.Error(), "\n")
	return r
}

func (r *Response[T]) Code(code int) Sender {
	r.code = code
	return r
}

func (r *Response[T]) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.code)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		return err
	}
	return nil
}
