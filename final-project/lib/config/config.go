package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DB  DB  `json:"db"`
	App App `json:"app"`
}

type DB struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (db DB) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", db.Username, db.Password, db.Host, db.Port, db.Name)
}

type App struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func Load(path string) (Config, error) {
	var conf Config
	if _, err := os.Stat(path); err != nil {
		return conf, err
	}

	confFile, err := os.Open(path)
	if err != nil {
		return conf, err
	}

	err = json.NewDecoder(confFile).Decode(&conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
