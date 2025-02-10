package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/MiikaMatias/portfolio/api/gen"
	"github.com/MiikaMatias/portfolio/internal/utils"
	"github.com/MiikaMatias/portfolio/server"
	"github.com/go-chi/chi/v5"
)

const (
	CONFIG_FILE_PATH = "./cmd/config.yaml"
)

func main() {
	config := utils.Config{}
	err := utils.CastConfig(CONFIG_FILE_PATH, &config)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	log.Printf("Config: %s\n", config.App.WorkingDir)

	r := chi.NewRouter()
	staticDir := filepath.Join(config.App.WorkingDir, "../templates")
	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir(staticDir))
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
	})
	log.Printf("Serving static files from: %s\n", staticDir)

	server := server.NewServer(config)
	h := gen.HandlerFromMux(server, r)

	portString := fmt.Sprintf(":%s", config.App.BackendPort)
	log.Printf("Server running on http://0.0.0.0%s\n", portString)
	s := &http.Server{
		Handler: h,
		Addr:    portString,
	}

	log.Fatal(s.ListenAndServe())
}
