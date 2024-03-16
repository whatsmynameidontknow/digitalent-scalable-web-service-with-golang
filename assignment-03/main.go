package main

import (
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const REFRESH_PERIOD_IN_SECONDS = 15 // page will refresh every min(15 - time since last generated, 15) seconds

//go:embed index.html
var indexDotHTML embed.FS

func main() {
	var jsonFilePath string
	flag.StringVar(&jsonFilePath, "json-file", "water-and-wind.json", "path to the json-file")
	flag.Parse()

	data, err := newData(jsonFilePath, REFRESH_PERIOD_IN_SECONDS)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer data.jsonFile.Close()

	tmpl, err := template.ParseFS(indexDotHTML, "*.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := http.NewServeMux()
	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			data.NextRefreshInSeconds = data.refreshPeriod - int(time.Since(data.lastGeneratedAt).Seconds())
			err := tmpl.ExecuteTemplate(w, "index", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.NotFound(w, r)
		}
	})

	server := new(http.Server)
	server.Handler = r
	server.Addr = fmt.Sprintf("0.0.0.0:%d", 8080)

	fmt.Printf("server berjalan di %s\n", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println(err)
		}

	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
	fmt.Println("shutting down server...")
	server.Shutdown(ctx)
}
