package board

import (
	"fmt"
	"strings"
)

type Coordinates struct {
	X uint8
	Y uint8
}

type PlayerAndMove struct {
	Move   Coordinates `json:"move"`
	Player string      `json:"player"`
}

type PlayerAndMoveEmbedded struct {
	Coordinates
	Player string
}

func (c Coordinates) Validate() error {
	if c.X > Length || c.Y > Length {
		return fmt.Errorf("outside of bounds")
	}

	return nil
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
