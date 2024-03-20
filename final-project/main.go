package main

import (
	"context"
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
)

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

	r := http.NewServeMux()

	{
		routes.InitUserRoutes(r, db, logger)
		routes.InitPhotoRoutes(r, db, logger)
		routes.InitCommentRoutes(r, db, logger)
		routes.InitSocialMediaRoutes(r, db, logger)
	}

	server := new(http.Server)
	server.Addr = fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	server.Handler = middleware.Recover(middleware.Logging(r))

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
