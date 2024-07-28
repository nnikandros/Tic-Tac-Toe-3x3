package test

import (
	"fmt"
	"testing"
	"tic-tac-toe/internal/board"
)

func TestBoard(t *testing.T) {

	// b := board.Board{}

	// fmt.Println(b)

	x := map[int]string{0: "kaaka", 1: "papiko"}

	if v, ok := x[3]; !ok {
		fmt.Println("NOT OK")
		fmt.Printf("vaue of string is %s", v)
	}

}

func TestBoard2(t *testing.T) {

	b := board.Board2{}

	fmt.Println(b)

	v, ok := b.Value(board.Coordinates{0, 0})

	fmt.Println(v)

	fmt.Println(ok)

}
