package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MiikaMatias/portfolio/api/gen"
	"github.com/MiikaMatias/portfolio/internal/llmclient"
	"github.com/MiikaMatias/portfolio/internal/utils"
)

const (
	TEMPLATE_DIR = "../templates/"
)

type Server struct {
	Config utils.Config
}

func NewServer(c utils.Config) Server {
	return Server{
		Config: c,
	}
}

func (s Server) GetEncodeString(w http.ResponseWriter, r *http.Request, params gen.GetEncodeStringParams) {
	log.Printf("Utilising model: %s\n", s.Config.App.EncodingModel)

	llmReply := llmclient.GetVectorEncoding(params.String, s.Config)

	resp := gen.EncodedVector{
		EncodedVector: llmReply,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) GetRagReply(w http.ResponseWriter, r *http.Request, params gen.GetRagReplyParams) {
	log.Printf("Utilising model: %s\n", s.Config.App.GeneratingModel)

	llmReply := llmclient.GetGeneratedResponse(params.Prompt, s.Config)

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

func (Server) Get(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, TEMPLATE_DIR+"index.html")
}
