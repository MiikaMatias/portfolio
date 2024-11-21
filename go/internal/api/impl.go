package api

import (
	"encoding/json"
	"net/http"

	"github.com/MiikaMatias/portfolio/internal/api/gen"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetAiResponse(w http.ResponseWriter, r *http.Request) {
	resp := gen.AiResponse{
		AiResponse: "generic response",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (Server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := "Pong"

	w.Write([]byte(resp))
}
