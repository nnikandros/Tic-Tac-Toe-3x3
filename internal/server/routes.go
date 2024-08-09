package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterRoutes() http.Handler {
	// mux := mux.NewRouter()
	r := chi.NewRouter()

	r.Get("/", s.HelloWorldHandler)
	return r

}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {

	resp := make(map[string]string)
	resp["message"] = "hello world"

	jsonResp, _ := json.Marshal(resp)

	// msg := []byte("hello world")
	w.Write(jsonResp)

}
