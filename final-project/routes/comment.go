package routes

import (
	"database/sql"
	"final-project/controller"
	"final-project/middleware"
	commentrepository "final-project/repository/comment"
	photorepository "final-project/repository/photo"
	commentservice "final-project/service/comment"
	"net/http"
)

func InitCommentRoutes(r *http.ServeMux, db *sql.DB) {
	commentRepo := commentrepository.New(db)
	photoRepo := photorepository.New(db)
	service := commentservice.New(commentRepo, photoRepo, db)
	controller := controller.NewCommentController(service)

	r.Handle("POST /comments", middleware.Auth(http.HandlerFunc(controller.Create)))
	r.Handle("GET /comments", middleware.Auth(http.HandlerFunc(controller.GetAll)))
	r.Handle("PUT /comments/{commentID}", middleware.Auth(http.HandlerFunc(controller.Update)))
	r.Handle("DELETE /comments/{commentID}", middleware.Auth(http.HandlerFunc(controller.Delete)))
}
