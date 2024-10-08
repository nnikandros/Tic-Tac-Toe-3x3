package board

import (
	"fmt"
	"strings"
)

const (
	Length = 3
	Width  = 3
)

var PlayerXMark = 1
var PlayerOMark = 0

type Board [Length][Width]*int

// Method that checks if at a given Coordinate a value has been set or not
// TODO: make b into a *Board2 for consistency
func (b Board) HasBeenSet(p Coordinates) (int, bool) {

	v := b[p.X][p.Y]

	if v == nil {
		return 0, false
	}

	return *v, true

}

// TODO parametrize the value that we puut. Mutates state. use for testing not for general use
func (b *Board) Set(p Coordinates) error {
	if value, ok := b.HasBeenSet(p); ok {
		return fmt.Errorf("at Coordidanates %v, value has been set already with value %v", p, value)
	}

	b[p.X][p.Y] = &PlayerXMark

	return nil

}

// Takes a board state and registers a move, a place in the grid together with the player.
// Returns a new Board state
func (b Board) RegisterMove(p PlayerAndMove) (Board, error) {
	if value, ok := b.HasBeenSet(p.Move); ok {
		return Board{}, fmt.Errorf("at Coordidanates %v, value has been set already with value %v", p, value)
	}

	playerSymbol := strings.ToLower(p.Player)
	if strings.Compare(playerSymbol, "x") == 0 {
		b[p.Move.X][p.Move.Y] = &PlayerXMark
	} else {
		b[p.Move.X][p.Move.Y] = &PlayerOMark

	}

	return b, nil

}
