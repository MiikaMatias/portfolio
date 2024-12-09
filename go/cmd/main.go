package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MiikaMatias/portfolio/api/gen"
	"github.com/MiikaMatias/portfolio/internal/llmclient"
	"github.com/MiikaMatias/portfolio/internal/redis"
	"github.com/MiikaMatias/portfolio/internal/utils"
	"github.com/go-chi/chi/v5"
)

const (
	CONFIG_FILE_PATH = "cmd/config.yaml"
)

type Config struct {
	App struct {
		Name        string `yaml:"name"`
		WorkingDir  string `yaml:"workingDir"`
		BackendPort string `yaml:"backendPort"`
	} `yaml:"app"`
}

func main() {
	config := Config{}
	err := utils.CastConfig(CONFIG_FILE_PATH, &config)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	log.Printf("Config: %s\n", config.App.WorkingDir)

	server := llmclient.NewServer()
	r := chi.NewMux()
	h := gen.HandlerFromMux(server, r)

	redisClient := redis.InitializeRedisClient()
	log.Printf("Initialization of client complete: %s", redisClient)

	portString := fmt.Sprintf("0.0.0.0:%s", config.App.BackendPort)
	log.Printf("Server running on http://%s\n", portString)
	s := &http.Server{
		Handler: h,
		Addr:    portString,
	}

	log.Fatal(s.ListenAndServe())
}
