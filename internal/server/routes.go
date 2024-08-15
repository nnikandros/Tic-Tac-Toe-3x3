package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tic-tac-toe/internal/board"

	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterRoutes() http.Handler {
	// mux := mux.NewRouter()
	r := chi.NewRouter()

	r.Get("/", s.HelloWorldHandler)
	r.Post("/posthandler", s.PostHandlerTest)
	return r

}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {

	resp := make(map[string]string)
	resp["message"] = "hello world"

	jsonResp, _ := json.Marshal(resp)

	// msg := []byte("hello world")
	w.Write(jsonResp)

}
func (s *Server) PostHandlerTest(w http.ResponseWriter, r *http.Request) {

	// resp := make(map[string]string)
	// resp["message"] = "hello world"

	// jsonResp, _ := json.Marshal(resp)

	// // msg := []byte("hello world")
	// w.Write(jsonResp)
	var t board.PlayerAndMoveNoRestraint
	Body, _ := io.ReadAll(r.Body)
	// decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	// err := decoder.Decode(&t)
	err := json.Unmarshal(Body, &t)

	if err != nil {
		fmt.Printf("error happend with %v", err)
	}

	NewPlayAndMove, err := t.Validate()

	if err != nil {
		fmt.Printf("error happend with %+v", err)
	}

	fmt.Println(string(Body))
	fmt.Printf("the body of the post request was %+v\n", t)

}

// type appHandler func(http.ResponseWriter, *http.Request) error
// func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//     if err := fn(w, r); err != nil {
//         http.Error(w, err.Error(), 500)
//     }
// }
