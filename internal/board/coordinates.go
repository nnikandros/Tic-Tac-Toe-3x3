package board

import "fmt"

type Coordinates struct {
	X uint8
	Y uint8
}

func (c Coordinates) Validate() error {
	if c.X > Length || c.X < 0 {
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
