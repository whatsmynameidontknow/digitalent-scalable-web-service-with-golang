package routes

import (
	"database/sql"
	"final-project/controller"
	"final-project/middleware"
	userrepository "final-project/repository/user"
	userservice "final-project/service/user"
	"log/slog"
	"net/http"
)

func InitUserRoutes(r *http.ServeMux, db *sql.DB, logger *slog.Logger) {
	userRepo := userrepository.New(db)
	userService := userservice.New(userRepo, logger)
	userController := controller.NewUserController(userService)

	r.Handle("PUT /users", middleware.Auth(http.HandlerFunc(userController.Update)))
	r.HandleFunc("POST /users/register", userController.Register)
	r.HandleFunc("POST /users/login", userController.Login)
	r.Handle("DELETE /users", middleware.Auth(http.HandlerFunc(userController.Delete)))
}
