package board

import (
	"fmt"
	"testing"
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

	b := Board{}

	fmt.Println(b)

	v, ok := b.HasBeenSet(Coordinates{0, 0})

	fmt.Println(v)

	fmt.Println(ok)

}

func TestBoard3(t *testing.T) {

	t.Run("testing for Set value with PlayerXMark ", func(t *testing.T) {
		b := Board{}

		b[0][0] = &PlayerXMark

		gotValue, gotOk := b.HasBeenSet(Coordinates{0, 0})
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
		b := Board{}

		b[1][1] = &PlayerOMark

		gotValue, gotOk := b.HasBeenSet(Coordinates{1, 1})
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
		b := Board{}

		_, gotOk := b.HasBeenSet(Coordinates{1, 1})

		if gotOk != false {
			t.Errorf("got %v but expteced false", gotOk)
		}

	})

	t.Run("testing for Set method return nil", func(t *testing.T) {

		b := Board{}

		err := b.Set(Coordinates{1, 1})

		if err != nil {
			t.Errorf("got %v but should be nil", err)
		}

	})
	t.Run("testing for Set method return not nill", func(t *testing.T) {

		b := Board{}

		b[1][1] = &PlayerOMark

		err := b.Set(Coordinates{1, 1})

		if err == nil {
			t.Errorf("got %v but should be not nil", err)
		}

	})
	t.Run("Testing for immutability", func(t *testing.T) {

		b := Board{}

		c := Coordinates{0, 0}
		p := "x"

		playerAndMove := PlayerAndMove{Move: c, Player: p}
		newState, _ := b.RegisterMove(playerAndMove)

		fmt.Printf("New Board state %v:", newState)
		fmt.Printf("\n \n Previous Board state %v:", b)
		// b2 := b

		// b[1][1] = &board.PlayerOMark

		// b2[0][0] = &board.PlayerXMark

		// fmt.Println(b2)

		// fmt.Println(b)

	})

	t.Run("Testing for mutation", func(t *testing.T) {

		// b := board.Board2{}

		// b2 := &b

		// // fmt.Println(b2)
		// b[1][1] = &board.PlayerOMark

		// b2[0][0] = &board.PlayerXMark

		// fmt.Println(b2)

		// fmt.Println(b)

	})

	t.Run("testing Register move with nil", func(t *testing.T) {
		b := Board{}

		playerAndMove := PlayerAndMove{Move: Coordinates{0, 0}, Player: "x"}

		_, err := b.RegisterMove(playerAndMove)

		if err != nil {
			t.Errorf("error should be nil here but got %v", err)
		}

	})
	t.Run("testing Register move with non nil", func(t *testing.T) {
		b := Board{}

		playerAndMove := PlayerAndMove{Move: Coordinates{1, 0}, Player: "x"}

		b.Set(Coordinates{0, 0})

		fmt.Println(b)

		_, err := b.RegisterMove(playerAndMove)

		if err != nil {
			t.Errorf("error should be non nil here but got %v", err)
		}

	})

}
