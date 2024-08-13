package board

import (
	"fmt"
	"strings"
)

type Coordinates struct {
	X uint8 `json:"x"`
	Y uint8 `json:"y"`
}

type CoordinatesNoRestraint struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PlayerAndMoveNoRestraint struct {
	Move   CoordinatesNoRestraint `json:"move"`
	Player string                 `json:"player"`
}

type PlayerAndMove struct {
	Move   Coordinates `json:"move"`
	Player string      `json:"player"`
}

func (c PlayerAndMoveNoRestraint) Validate() (PlayerAndMove, error) {
	x := c.Move.X
	y := c.Move.Y
	u := strings.ToLower(c.Player)

	if x < 0 || x > Length {
		return PlayerAndMove{}, fmt.Errorf("out of bounds")
	}

	if y < 0 || y > Length {
		return PlayerAndMove{}, fmt.Errorf("out of bounds")
	}

	if strings.Compare(u, "x") != 0 && strings.Compare(u, "o") != 0 {
		return PlayerAndMove{}, fmt.Errorf("error")

	}

	return PlayerAndMove{Move: Coordinates{uint8(x), uint8(y)}, Player: u}, nil

}

func NewCoordinatesFromRequest(x, y int) (Coordinates, error) {

	if x < 0 || x > Length {
		return Coordinates{}, fmt.Errorf("out of bounds")
	}

	if y < 0 || y > Length {
		return Coordinates{}, fmt.Errorf("out of bounds")
	}

	return Coordinates{X: uint8(x), Y: uint8(y)}, nil
}

func ValidatePlayer(p PlayerAndMove) error {
	u := strings.ToLower(p.Player)

	if strings.Compare(u, "x") != 0 && strings.Compare(u, "o") != 0 {
		return fmt.Errorf("error")

	}

	return nil

}
