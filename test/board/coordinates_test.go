package test

import (
	"reflect"
	"testing"
	"tic-tac-toe/internal/board"
)

func TestNewCoordidantesFromReques(t *testing.T) {

	t.Run("correct things", func(t *testing.T) {

		x := 0
		y := 1

		got, _ := board.NewCoordinatesFromRequest(x, y)

		expected := board.Coordinates{0, 1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got  %v but expected %v", got, expected)
		}
	})

	t.Run("correct things 2", func(t *testing.T) {

		x := 2
		y := 1

		got, _ := board.NewCoordinatesFromRequest(x, y)

		expected := board.Coordinates{2, 1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got  %v but expected %v", got, expected)
		}

	})

	t.Run("wrong things", func(t *testing.T) {

		x := -2
		y := 1

		got, err := board.NewCoordinatesFromRequest(x, y)

		expected := board.Coordinates{}

		if !reflect.DeepEqual(got, expected) {
		}

		if err == nil {
			t.Errorf("error should be non nil ")
		}

	})

	t.Run("wrong things 2", func(t *testing.T) {

		x := 5
		y := 0

		got, err := board.NewCoordinatesFromRequest(x, y)

		expected := board.Coordinates{}

		if !reflect.DeepEqual(got, expected) {
		}

		if err == nil {
			t.Errorf("error should be non nil ")
		}

	})

}
