package main

import (
	"encoding/json"
	"math/rand/v2"
	"os"
	"time"
)

type WaterAndWind struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Data struct {
	Status               WaterAndWind `json:"status"`
	NextRefreshInSeconds int          `json:"-"`
	jsonFile             *os.File
	lastGeneratedAt      time.Time
	refreshPeriod        int
}

func newData(filePath string, refreshPeriod int) (*Data, error) {
	data := new(Data)
	err := data.load(filePath)
	if err != nil {
		return nil, err
	}

	data.refreshPeriod = refreshPeriod
	data.lastGeneratedAt = time.Now()
	go data.generateRandom()

	return data, nil
}

func (w *Data) load(filePath string) error {
	jsonFile, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	w.jsonFile = jsonFile
	return json.NewDecoder(w.jsonFile).Decode(w)
}

func (w *Data) write() error {
	w.jsonFile.Truncate(0)
	w.jsonFile.Seek(0, 0)
	return json.NewEncoder(w.jsonFile).Encode(w)
}

// generate random data every REFRESH_PERIOD_IN_SECONDS seconds
// the data is generated in the backend instead of every time client refreshes the page
// so that the data is consistent across all clients
func (w *Data) generateRandom() {
	ticker := time.NewTicker(time.Duration(w.refreshPeriod) * time.Second)
	for range ticker.C {
		w.Status.Water = 1 + rand.IntN(100)
		w.Status.Wind = 1 + rand.IntN(100)
		w.write()
		w.lastGeneratedAt = time.Now()
	}
}
