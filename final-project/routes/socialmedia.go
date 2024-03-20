package routes

import (
	"database/sql"
	"final-project/controller"
	"final-project/middleware"
	socialmediarepository "final-project/repository/socialmedia"
	socialmediaservice "final-project/service/socialmedia"
	"log/slog"
	"net/http"
)

func InitSocialMediaRoutes(r *http.ServeMux, db *sql.DB, logger *slog.Logger) {
	socialMediaRepo := socialmediarepository.New(db)
	socialMediaService := socialmediaservice.New(socialMediaRepo, db, logger)
	socialMediaController := controller.NewSocialMediaController(socialMediaService)

	r.Handle("POST /socialmedias", middleware.Auth(http.HandlerFunc(socialMediaController.Create)))
	r.Handle("GET /socialmedias", middleware.Auth(http.HandlerFunc(socialMediaController.GetAll)))
	r.Handle("PUT /socialmedias/{socialMediaID}", middleware.Auth(http.HandlerFunc(socialMediaController.Update)))
	r.Handle("DELETE /socialmedias/{socialMediaID}", middleware.Auth(http.HandlerFunc(socialMediaController.Delete)))
}
