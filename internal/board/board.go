package board

import "fmt"

const Length = 3
const Width = 3

var PlayerXMark = 1
var PlayerOMark = 0

// type Board [Length][Width]int8

// func (b Board) HasBeenset(p Coordinates) (int8, bool) {
// 	v := b[p.X][p.Y]
// 	if v == 0 {
// 		return 0, false
// 	}

// 	return v, true

// }

type Board2 [Length][Width]*int

func (b Board2) HasBeenset(p Coordinates) (int, bool) {

	v := b[p.X][p.Y]

	if v == nil {
		return 0, false
	}

	return *v, true

}

func (b *Board2) Set(p Coordinates) error {
	if value, ok := b.HasBeenset(p); ok {
		return fmt.Errorf("at Coordidanates %v, value has been set already with value %v", p, value)
	}

	b[p.X][p.Y] = &PlayerXMark

	return nil

}
