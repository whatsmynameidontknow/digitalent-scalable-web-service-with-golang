package main

import (
	"assignment-02/config"
	"assignment-02/controller"
	"assignment-02/docs"
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

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Hacktiv8-Golang assignment-02
// @version 1.0
// @description submission for assignment-02
// @BasePath /
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

	serverAddr := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	docs.SwaggerInfo.Host = serverAddr

	orderRepo := repository.NewOrderRepository(db)
	itemRepo := repository.NewItemRepository()
	orderService := service.NewOrderService(db, orderRepo, itemRepo, logger)
	orderController := controller.NewOrderController(orderService)

	api := http.NewServeMux()
	{
		api.HandleFunc("POST /orders", orderController.Create)
		api.HandleFunc("GET /orders", orderController.GetAll)
		api.HandleFunc("GET /orders/{orderId}", orderController.GetByID)
		api.HandleFunc("DELETE /orders/{orderId}", orderController.Delete)
		api.HandleFunc("PUT /orders/{orderId}", orderController.Update)
	}

	r := http.NewServeMux()
	r.Handle("/", middleware.Logging(api))
	r.HandleFunc("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", serverAddr)),
	))

	srv := new(http.Server)
	srv.Handler = middleware.Recover(r)
	srv.Addr = serverAddr

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
