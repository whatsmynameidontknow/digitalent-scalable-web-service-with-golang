package main

import (
	"context"
	"final-project/lib/config"
	"final-project/lib/database"
	"final-project/routes"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.Load("config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := database.New(conf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := http.NewServeMux()

	{
		routes.InitUserRoutes(r, db)
	}

	server := new(http.Server)
	server.Addr = fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	server.Handler = r

	fmt.Printf("Server started at %s\n", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	<-ctx.Done()

	fmt.Println("Shutting down server...")
	server.Shutdown(ctx)
}
