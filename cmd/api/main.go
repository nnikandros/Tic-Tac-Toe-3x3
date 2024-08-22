package main

import (
	"fmt"
	directoryregistry "tic-tac-toe/internal/directory_registry"
	"tic-tac-toe/internal/server"
)

func main() {

	err := directoryregistry.CreateDirs()

	if err != nil {
		fmt.Printf("probably games dir already exists, %v\n", err)
	}

	server := server.NewServer()

	err = server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
