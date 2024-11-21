package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MiikaMatias/portfolio/internal/api"
	"github.com/MiikaMatias/portfolio/internal/api/gen"
	"github.com/go-chi/chi/v5"
	"gopkg.in/yaml.v3"
)

const (
	CONFIG_FILE_PATH = "cmd/config.yaml"
)

type Config struct {
	App struct {
		Name       string `yaml:"name"`
		WorkingDir string `yaml:"workingDir"`
	} `yaml:"app"`
}

func loadConfig() (*Config, error) {
	file, err := os.Open(CONFIG_FILE_PATH)
	if err != nil {
		return nil, err
	}

	var config Config
	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Printf("Config: %s", config.App.WorkingDir)

	server := api.NewServer()
	r := chi.NewMux()
	h := gen.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
