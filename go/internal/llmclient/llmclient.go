package llmclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MiikaMatias/portfolio/internal/utils"
)

const (
	CONFIG_FILE_PATH = "internal/llmclient/config.yaml"
)

type VectorEmbedding struct {
	Embedding []float32 `json:"embedding"`
}

type LlmReply struct {
	Response string `json:"response"`
}

func GetVectorEncoding(prompt string, config utils.Config) []float32 {
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
	defer res.Body.Close()

	vectorReply := &VectorEmbedding{}
	derr := json.NewDecoder(res.Body).Decode(vectorReply)
	if derr != nil {
		log.Fatalf("Error decoding Ollama reply: %s", err)
	}

	return vectorReply.Embedding
}

func GetGeneratedResponse(prompt string, config utils.Config) string {
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
	defer res.Body.Close()

	llmReply := &LlmReply{}
	derr := json.NewDecoder(res.Body).Decode(llmReply)
	if derr != nil {
		log.Fatalf("Error decoding Ollama reply: %s", err)
	}
	log.Printf("Decoded response: %v", llmReply.Response)

	return llmReply.Response
}
