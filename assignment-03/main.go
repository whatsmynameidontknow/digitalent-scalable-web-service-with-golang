package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"math/rand/v2"
	"net/http"
	"os"
	"time"
)

const REFRESH_PERIOD_IN_SECONDS = 15 // page will refresh every min(15 - time since last generated, 15) seconds

//go:embed index.html
var indexDotHTML embed.FS

var (
	data     = new(Data)
	jsonFile string
)

func init() {
	flag.StringVar(&jsonFile, "json-file", "water-and-wind.json", "path to the json-file")
	flag.Parse()
	err := data.load(jsonFile)
	if err != nil {
		panic(err)
	}
	data.NextRefreshInSeconds = REFRESH_PERIOD_IN_SECONDS
	data.lastGeneratedAt = time.Now()
}

func main() {
	server := http.NewServeMux()
	tmpl, err := template.ParseFS(indexDotHTML, "*.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func() {
		// generate random data every REFRESH_PERIOD_IN_SECONDS seconds
		// the data is generated in the backend instead of every time client refreshes the page
		// so that the data is consistent across all clients
		ticker := time.NewTicker(REFRESH_PERIOD_IN_SECONDS * time.Second)
		for range ticker.C {
			data.generateRandom()
			data.lastGeneratedAt = time.Now()
		}
	}()

	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			data.NextRefreshInSeconds = REFRESH_PERIOD_IN_SECONDS - int(time.Since(data.lastGeneratedAt).Seconds())
			err := tmpl.ExecuteTemplate(w, "index", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.NotFound(w, r)
		}
	})

	port := 8080
	fmt.Printf("server jalan di port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), server)
}

type WaterAndWind struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
type Data struct {
	Status               WaterAndWind `json:"status"`
	NextRefreshInSeconds int          `json:"-"`
	filePath             string
	lastGeneratedAt      time.Time
}

func (w *Data) load(filePath string) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	w.filePath = filePath
	return json.NewDecoder(jsonFile).Decode(&w)
}

func (w *Data) write() error {
	jsonFile, err := os.OpenFile(w.filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	return json.NewEncoder(jsonFile).Encode(w)
}

func (w *Data) generateRandom() {
	w.Status.Water = 1 + rand.IntN(100)
	w.Status.Wind = 1 + rand.IntN(100)
	w.write()
}
