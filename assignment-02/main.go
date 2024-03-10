package main

import (
	"assignment-02/config"
	"assignment-02/controller"
	"assignment-02/lib/database"
	"assignment-02/lib/logging"
	"assignment-02/middleware"
	"assignment-02/repository"
	"assignment-02/service"
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "config.json", "path to JSON config file")
	flag.Parse()

	conf, err := config.Load(configFilePath)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	db, err := database.New(conf.DB)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	logger := logging.New(os.Stderr)

	orderRepo := repository.NewOrderRepository(db)
	itemRepo := repository.NewItemRepository()
	orderService := service.NewOrderService(db, orderRepo, itemRepo, logger)
	orderController := controller.NewOrderController(orderService)

	r := http.NewServeMux()

	{
		r.HandleFunc("POST /orders", orderController.Create)
		r.HandleFunc("GET /orders", orderController.GetAll)
		r.HandleFunc("GET /orders/{orderId}", orderController.GetByID)
		r.HandleFunc("DELETE /orders/{orderId}", orderController.Delete)
		r.HandleFunc("PUT /orders/{orderId}", orderController.Update)
	}

	srv := new(http.Server)
	srv.Handler = middleware.Recover(middleware.Logging(r))
	srv.Addr = fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)

	fmt.Printf("server is running at %s\n", srv.Addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			os.Stderr.WriteString(err.Error())
			os.Exit(1)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	<-ctx.Done()
	if err := srv.Shutdown(ctx); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	fmt.Println("server shut down!")
}
