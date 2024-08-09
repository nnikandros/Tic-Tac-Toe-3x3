package server

import "net/http"

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("GET /", s.HelloWorldHandler)
	return mux

}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {

	msg := []byte("hello world")
	w.Write(msg)

}
