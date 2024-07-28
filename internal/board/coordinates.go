package board

import (
	"fmt"
)

type Coordinates struct {
	X uint8
	Y uint8
}

func (c Coordinates) Validate() error {
	if c.X > Length || c.Y > Length {
		return fmt.Errorf("outside of bounds")
	}

	return nil
}

func NewCoordinatesFromRequest(x, y int) (Coordinates, error) {

	// max := math.Max(float64(x), float64(y))

	// min := math.Min(float64(x), float64(y))

	if x < 0 || x > Length {
		return Coordinates{}, fmt.Errorf("out of bounds")
	}

	if y < 0 || y > Length {
		return Coordinates{}, fmt.Errorf("out of bounds")
	}

	// fmt.Println(max)
	// if max < 0 || min > Length {
	// 	return Coordinates{}, fmt.Errorf("out of bounds")
	// }

	return Coordinates{X: uint8(x), Y: uint8(y)}, nil
}
