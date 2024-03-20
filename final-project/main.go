package main

import (
	"context"
	"final-project/docs"
	"final-project/helper"
	"final-project/lib/config"
	"final-project/lib/database"
	"final-project/lib/logging"
	"final-project/middleware"
	"final-project/routes"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Hacktiv8-Golang final-project
// @version 1.0
// @description submission for final-project
// @BasePath /
func main() {
	logger := logging.New(os.Stdout)
	conf, err := config.Load("config.json")
	if err != nil {
		logger.Error(err.Error(), "cause", "config.Load")
		os.Exit(1)
	}

	helper.JWTSecret = []byte(conf.App.JWTSecret)
	helper.JWTExpiresIn = helper.GetJWTExpiresIn(conf.App.JWTExpiresIn, time.Hour)

	db, err := database.New(conf.DB)
	if err != nil {
		logger.Error(err.Error(), "cause", "database.New")
		os.Exit(1)
	}
	defer db.Close()

	api := http.NewServeMux()

	{
		routes.InitUserRoutes(api, db, logger)
		routes.InitPhotoRoutes(api, db, logger)
		routes.InitCommentRoutes(api, db, logger)
		routes.InitSocialMediaRoutes(api, db, logger)
	}

	r := http.NewServeMux()
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)

	{
		r.Handle("/", middleware.Logging(api))
		r.HandleFunc("GET /swagger/", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://%s:%d/swagger/doc.json", conf.App.Host, conf.App.Port)),
		))
	}

	server := new(http.Server)
	server.Addr = fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	server.Handler = middleware.Recover(r)

	logger.Info("Starting server...", "addr", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(err.Error(), "cause", "server.ListenAndServe")
			os.Exit(1)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	<-ctx.Done()

	logger.Info("Shutting down server...", "addr", server.Addr)
	err = server.Shutdown(ctx)
	if err != nil {
		logger.Error(err.Error(), "cause", "server.Shutdown")
	}
}
