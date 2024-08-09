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

	v, ok := b.HasBeenSet(board.Coordinates{0, 0})

	fmt.Println(v)

	fmt.Println(ok)

}

func TestBoard3(t *testing.T) {

	t.Run("testing for Set value with PlayerXMark ", func(t *testing.T) {
		b := board.Board2{}

		b[0][0] = &board.PlayerXMark

		gotValue, gotOk := b.HasBeenSet(board.Coordinates{0, 0})
		expectedOk := true
		expectedValue := 1

		if expectedOk != gotOk {
			t.Errorf("expected %v but got %v", expectedOk, gotOk)
		}

		if expectedValue != gotValue {
			t.Errorf("expected %v but got %v", expectedValue, expectedValue)
		}
	})

	t.Run("testing for Set value with PlayerOMark ", func(t *testing.T) {
		b := board.Board2{}

		b[1][1] = &board.PlayerOMark

		gotValue, gotOk := b.HasBeenSet(board.Coordinates{1, 1})
		expectedOk := true
		expectedValue := 0

		if expectedOk != gotOk {
			t.Errorf("expected %v but got  %v", expectedOk, gotOk)
		}

		if expectedValue != gotValue {
			t.Errorf("expected %v but got %v", expectedValue, gotValue)
		}
	})

	t.Run("testing for Set value with nil pointer ", func(t *testing.T) {
		b := board.Board2{}

		_, gotOk := b.HasBeenSet(board.Coordinates{1, 1})

		if gotOk != false {
			t.Errorf("got %v but expteced false", gotOk)
		}

	})

	t.Run("testing for Set method return nil", func(t *testing.T) {

		b := board.Board2{}

		err := b.Set(board.Coordinates{1, 1})

		if err != nil {
			t.Errorf("got %v but should be nil", err)
		}

	})
	t.Run("testing for Set method return not nill", func(t *testing.T) {

		b := board.Board2{}

		b[1][1] = &board.PlayerOMark

		err := b.Set(board.Coordinates{1, 1})

		if err == nil {
			t.Errorf("got %v but should be not nil", err)
		}

	})
	t.Run("Testing for immutability", func(t *testing.T) {

		// b := board.Board2{}

		// b2 := b

		// b[1][1] = &board.PlayerOMark

		// b2[0][0] = &board.PlayerXMark

		// fmt.Println(b2)

		// fmt.Println(b)

	})

	t.Run("Testing for mutation", func(t *testing.T) {

		b := board.Board2{}

		b2 := &b

		// fmt.Println(b2)
		b[1][1] = &board.PlayerOMark

		b2[0][0] = &board.PlayerXMark

		fmt.Println(b2)

		fmt.Println(b)

	})

}
