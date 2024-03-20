package routes

import (
	"database/sql"
	"final-project/controller"
	"final-project/middleware"
	socialmediarepository "final-project/repository/socialmedia"
	socialmediaservice "final-project/service/socialmedia"
	"net/http"
)

func InitSocialMediaRoutes(r *http.ServeMux, db *sql.DB) {
	socialMediaRepo := socialmediarepository.New(db)
	socialMediaService := socialmediaservice.New(socialMediaRepo, db)
	socialMediaController := controller.NewSocialMediaController(socialMediaService)

	r.Handle("POST /socialmedias", middleware.Auth(http.HandlerFunc(socialMediaController.Create)))
	r.Handle("GET /socialmedias", middleware.Auth(http.HandlerFunc(socialMediaController.GetAll)))
	r.Handle("PUT /socialmedias/{socialMediaID}", middleware.Auth(http.HandlerFunc(socialMediaController.Update)))
	r.Handle("DELETE /socialmedias/{socialMediaID}", middleware.Auth(http.HandlerFunc(socialMediaController.Delete)))
}
