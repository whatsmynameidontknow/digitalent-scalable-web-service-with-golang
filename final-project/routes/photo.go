package routes

import (
	"database/sql"
	"final-project/controller"
	"final-project/middleware"
	photorepository "final-project/repository/photo"
	photoservice "final-project/service/photo"
	"net/http"
)

func InitPhotoRoutes(r *http.ServeMux, db *sql.DB) {
	repo := photorepository.New(db)
	service := photoservice.New(repo, db)
	controller := controller.NewPhotoController(service)

	r.Handle("POST /photos", middleware.Auth(http.HandlerFunc(controller.Create)))
	r.Handle("GET /photos", middleware.Auth(http.HandlerFunc(controller.GetAll)))
	r.Handle("PUT /photos/{photoID}", middleware.Auth(http.HandlerFunc(controller.Update)))
	r.Handle("DELETE /photos/{photoID}", middleware.Auth(http.HandlerFunc(controller.Delete)))
}
