package board

const Length = 3
const Width = 3

type Board [Length][Width]int8

func (b Board) HasBeenset(p Coordinates) (int8, bool) {
	v := b[p.X][p.Y]
	if v == 0 {
		return 0, false
	}

	return v, true

}
