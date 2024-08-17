package board

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCoordidantesFromReques(t *testing.T) {

	t.Run("correct things", func(t *testing.T) {

		x := 0
		y := 1

		got, _ := NewCoordinatesFromRequest(x, y)

		expected := Coordinates{0, 1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got  %v but expected %v", got, expected)
		}
	})

	t.Run("correct things 2", func(t *testing.T) {

		x := 2
		y := 1

		got, _ := NewCoordinatesFromRequest(x, y)

		expected := Coordinates{2, 1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got  %v but expected %v", got, expected)
		}

	})

	t.Run("wrong things", func(t *testing.T) {

		x := -2
		y := 1

		got, err := NewCoordinatesFromRequest(x, y)

		expected := Coordinates{}

		if !reflect.DeepEqual(got, expected) {
		}

		if err == nil {
			t.Errorf("error should be non nil ")
		}

	})

	t.Run("wrong things 2", func(t *testing.T) {

		x := 5
		y := 0

		got, err := NewCoordinatesFromRequest(x, y)

		expected := Coordinates{}

		if !reflect.DeepEqual(got, expected) {
		}

		if err == nil {
			t.Errorf("error should be non nil ")
		}

	})
	// t.Run("saving board state to gob", func(t *testing.T) {
	// 	var buffer bytes.Buffer
	// 	encoder := gob.NewEncoder(&buffer)
	// 	b := board.Board{}

	// 	err := encoder.Encode(b)
	// 	if err != nil {
	// 		fmt.Println("Error encoding:", err)
	// 	}

	// 	file, err := os.Create("person.gob")

	// 	if err != nil {
	// 		fmt.Println("Error creating file:", err)
	// 	}
	// 	defer file.Close()

	// 	_, err = buffer.WriteTo(file)
	// 	if err != nil {
	// 		fmt.Println("Error writing to file:", err)
	// 	}

	// 	fmt.Println(uuid.NewString())

	// })
	t.Run("saving board state to gob", func(t *testing.T) {

		var b Board

		c := Coordinates{X: 0, Y: 0}

		p := PlayerAndMove{Move: c, Player: "X"}

		newBoard, _ := b.RegisterMove(p)
		var h []Board

		h = append(h, newBoard)

		fmt.Printf("%v\n\n", h)

		c1 := Coordinates{X: 1, Y: 0}

		p1 := PlayerAndMove{Move: c1, Player: "O"}

		newBoard2, _ := newBoard.RegisterMove(p1)

		h = append(h, newBoard2)

		fmt.Printf("%+v", h)
	})

}
