package board

const Length = 3
const Width = 3

var PlayerXMark = 1
var PlayerOMark = 0

type Board [Length][Width]int8

func (b Board) HasBeenset(p Coordinates) (int8, bool) {
	v := b[p.X][p.Y]
	if v == 0 {
		return 0, false
	}

	return v, true

}

type Board2 [Length][Width]*int

func (b Board2) Value(p Coordinates) (int, bool) {

	v := b[p.X][p.Y]

	if v == nil {
		return 0, false
	}

	return *v, true

}

func (b *Board2) Set(p Coordinates) {
	if value, ok := b.Value(p); ok {
	}

	

}
