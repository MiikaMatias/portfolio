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
		EncodingUrl     string `yaml:"encodingUrl"`
		EncodingModel   string `yaml:"encodingModel"`
		GeneratingUrl   string `yaml:"generatingUrl"`
		GeneratingModel string `yaml:"generatingModel"`
	} `yaml:"app"`
}

type VectorEmbedding struct {
	Embedding []float32 `json:"embedding"`
}

type LlmReply struct {
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

func getVectorEncoding(prompt string, config Config) []float32 {
	body := []byte(fmt.Sprintf(`{
		"model": "%s",
		"prompt": "%s"
	}`, config.App.EncodingModel, prompt))

	r, err := http.NewRequest("POST", config.App.EncodingUrl, bytes.NewBuffer(body))
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

	vectorReply := &VectorEmbedding{}
	derr := json.NewDecoder(res.Body).Decode(vectorReply)
	if derr != nil {
		log.Fatalf("Error decoding Ollama reply: %s", err)
	}

	return vectorReply.Embedding
}

func getGeneratedResponse(prompt string, config Config) string {
	body := []byte(fmt.Sprintf(`{
		"model": "%s",
		"prompt": "%s",
		"stream": false
	}`, config.App.GeneratingModel, prompt))

	r, err := http.NewRequest("POST", config.App.GeneratingUrl, bytes.NewBuffer(body))
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

	llmReply := &LlmReply{}
	derr := json.NewDecoder(res.Body).Decode(llmReply)
	if derr != nil {
		log.Fatalf("Error decoding Ollama reply: %s", err)
	}
	log.Printf("Decoded response: %v", llmReply.Response)

	return llmReply.Response
}

func (Server) GetEncodeString(w http.ResponseWriter, r *http.Request) {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	log.Printf("Utilising model: %s\n", config.App.EncodingModel)

	llmReply := getVectorEncoding("hello", *config)

	log.Printf("Reply decoded:", llmReply)

	resp := gen.EncodedVector{
		EncodedVector: llmReply,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (Server) GetRagReply(w http.ResponseWriter, r *http.Request) {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	log.Printf("Utilising model: %s\n", config.App.GeneratingModel)

	llmReply := getGeneratedResponse("hello", *config)

	log.Printf("Reply decoded:", llmReply)

	resp := gen.RagReply{
		RagReply: llmReply,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (Server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := "Pong"

	w.Write([]byte(resp))
}
