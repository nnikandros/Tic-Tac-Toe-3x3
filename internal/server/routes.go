package server

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"tic-tac-toe/internal/board"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s *Server) RegisterRoutes() http.Handler {
	// mux := mux.NewRouter()
	r := chi.NewRouter()

	r.Get("/", s.HelloWorldHandler)
	r.Get("/newgame", s.NewGame)
	r.Post("/posthandler", s.PostHandlerTest)
	return r

}

func (s *Server) NewGame(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["new_game"] = "created"

	gameId := uuid.NewString()
	resp["Id"] = gameId

	var buffer bytes.Buffer
	var h []board.Board
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(h)
	if err != nil {
		//   do something here
	}

	nameOfFile := strings.Join([]string{gameId, "gob"}, ".")

	file, err := os.Create(nameOfFile)

	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()

	_, err = buffer.WriteTo(file)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {

	resp := make(map[string]string)
	resp["message"] = "hello world"

	jsonResp, _ := json.Marshal(resp)

	// msg := []byte("hello world")
	w.Write(jsonResp)

}
func (s *Server) PostHandlerTest(w http.ResponseWriter, r *http.Request) {
	NewPlayAndMove, err := NewPlayAndMove(r)

}

func NewPlayAndMove(r *http.Request) (board.PlayerAndMove, error) {
	d := board.PlayerAndMove{}
	var t board.PlayerAndMoveNoRestraint
	Body, _ := io.ReadAll(r.Body)
	// decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	// err := decoder.Decode(&t)
	err := json.Unmarshal(Body, &t)

	if err != nil {
		return d, fmt.Errorf("error happend with %v", err)
	}

	NewPlayAndMove, err := t.Validate()

	fmt.Println(NewPlayAndMove)

	if err != nil {
		return d, fmt.Errorf("error happend with %+v", err)
	}

	fmt.Println(string(Body))
	fmt.Printf("the body of the post request was %+v\n", t)

	return NewPlayAndMove, nil

}

// type appHandler func(http.ResponseWriter, *http.Request) error
// func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//     if err := fn(w, r); err != nil {
//         http.Error(w, err.Error(), 500)
//     }
// }

// resp := make(map[string]string)
// resp["message"] = "hello world"

// jsonResp, _ := json.Marshal(resp)

// // msg := []byte("hello world")
// w.Write(jsonResp)
// var t board.PlayerAndMoveNoRestraint
// Body, _ := io.ReadAll(r.Body)
// // decoder := json.NewDecoder(r.Body)
// defer r.Body.Close()
// // err := decoder.Decode(&t)
// err := json.Unmarshal(Body, &t)

// if err != nil {
// 	fmt.Printf("error happend with %v", err)
// }

// NewPlayAndMove, err := t.Validate()

// if err != nil {
// 	fmt.Printf("error happend with %+v", err)
// }

// fmt.Println(string(Body))
// fmt.Printf("the body of the post request was %+v\n", t)
