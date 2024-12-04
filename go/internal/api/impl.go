package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MiikaMatias/portfolio/internal/api/gen"
	"gopkg.in/yaml.v2"
)

const (
	CONFIG_FILE_PATH = "internal/api/config.yaml"
)

type Config struct {
	App struct {
		OllamaUrl        string `yaml:"ollamaUrl"`
		OllamaModelToUse string `yaml:"ollamaModelToUse"`
	} `yaml:"app"`
}

type OllamaReply struct {
	Response string `json:"response"`
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

type Server struct{}

func NewServer() Server {
	return Server{}
}

func fetchLlmReply(config Config) string {
	body := []byte(fmt.Sprintf(`{
		"model": "%s",
		"prompt": "Why is the sky blue?",
		"stream": false
	}`, config.App.OllamaModelToUse))

	r, err := http.NewRequest("POST", config.App.OllamaUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Error constructing request: %s", err)
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatalf("Error sending POST request: %s", err)
	}
	log.Printf("Received response: %s", res)
	defer res.Body.Close()

	llmReply := &OllamaReply{}
	derr := json.NewDecoder(res.Body).Decode(llmReply)
	if derr != nil {
		log.Fatalf("Error decoding Ollama reply: %s", err)
	}

	return llmReply.Response
}

func (Server) GetAiResponse(w http.ResponseWriter, r *http.Request) {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	log.Printf("Utilising model: %s\n", config.App.OllamaModelToUse)

	llmReply := fetchLlmReply(*config)

	log.Printf("Reply decoded:", llmReply)

	resp := gen.AiResponse{
		AiResponse: string(llmReply),
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (Server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := "Pong"

	w.Write([]byte(resp))
}
