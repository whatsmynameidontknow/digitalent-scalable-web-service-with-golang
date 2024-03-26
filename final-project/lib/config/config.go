package config

import (
	"encoding/json"
	"final-project/helper"
	"fmt"
	"net/url"
	"os"
	"regexp"
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
	Host         string `json:"host"`
	Port         uint   `json:"port"`
	JWTSecret    string `json:"jwt_secret"`
	JWTExpiresIn string `json:"jwt_expires_in"`
	BasePath     string `json:"base_path"`
}

func (app App) isValidBasePath() bool {
	url, err := url.Parse(app.BasePath)

	return err == nil && url.Path == app.BasePath && regexp.MustCompile(`^\/([a-zA-Z0-9-]+\/)*$|^\/$`).MatchString(app.BasePath)
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

	if !conf.App.isValidBasePath() {
		return conf, helper.ErrInvalidBasePath
	}

	return conf, nil
}
