package test

import (
	"fmt"
	"testing"
	"tic-tac-toe/internal/board"
)

func TestHasBeenSet(t *testing.T) {

	t.Run("first test", func(t *testing.T) {

		b := board.Board{}
		p := board.Coordinates{0, 0}
		u, ok := b.HasBeenset(p)

		fmt.Println(ok)
		fmt.Println(u)

	})

}
