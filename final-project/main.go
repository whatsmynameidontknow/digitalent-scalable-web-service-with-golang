package main

import (
	"context"
	"final-project/helper"
	"final-project/lib/config"
	"final-project/lib/database"
	"final-project/lib/logging"
	"final-project/middleware"
	"final-project/routes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "final-project/docs"

	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Hacktiv8-Golang final-project
// @version 1.0
// @description submission for final-project
// @BasePath /api/v1
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
// @description Bearer token for authentication. Format: Bearer {token}
func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "json-config", "config.json", "path to json config file")
	flag.Parse()

	logger := logging.New(os.Stdout)
	conf, err := config.Load(configFilePath)
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

	{
		r.Handle("/api/v1/", middleware.Logging(http.StripPrefix("/api/v1", api)))
		r.HandleFunc("GET /swagger/", httpSwagger.Handler(
			httpSwagger.URL("/swagger/doc.json"),
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
