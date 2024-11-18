package api

import (
	"encoding/json"
	"net/http"

	"github.com/MiikaMatias/portfolio/api/spec"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := spec.Pong{
		Ping: "pong",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
